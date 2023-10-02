package builder

import (
	"prism/internal/model"
	"slices"
)

var (
	singleType  = "singleBlock"
	unnamedType = "unnamedBlock"
	namedType   = "namedBlock"
)

type Changes struct{}

func NewChanges() *Changes {
	return &Changes{}
}

// Making changes to the configuration file.
// Adds parameters and blocks specified in additional files,
// parameters from flags and environment variables.
func (s *Changes) SetChanges(
	config *model.TemplateBlock,
	changes *model.Changes,
) error {
	blockChanges := model.BlockChanges{
		ProjectDirPath: changes.ProjectDirPath,
		Release:        changes.Release,
		Namespace:      changes.Namespace,
		File:           model.TemplateBlock{},
		Chart:          changes.Chart,
	}

	if len(changes.Files) > 0 {
		for _, file := range changes.Files {
			blockChanges.File = file
			blockChanges.Release = ""
			job(config, &blockChanges)
		}

		if changes.Release != "" {
			blockChanges.File = model.TemplateBlock{}
			blockChanges.Release = changes.Release
			job(config, &blockChanges)
		}

		return nil
	}

	job(config, &blockChanges)

	return nil
}

// Checks for the presence of blocks
// in the specified file that can be specified only once.
// If the block specified in the file is in the configuration, it is ignored.
// Otherwise it will be added.
func checkSingleBlocks(block, file *model.TemplateBlock, blockName []string) {
	var haveBlock []string

	for _, fileBlock := range file.Block {
		if slices.Contains(blockName, fileBlock.BlockName) {
			for _, item := range block.Block {
				if item.BlockName == fileBlock.BlockName {
					haveBlock = append(haveBlock, fileBlock.BlockName)
				}
			}
		}
	}

	for _, item := range file.Block {
		if slices.Contains(blockName, item.BlockName) {
			if !slices.Contains(haveBlock, item.BlockName) {
				block.Block = append(block.Block, item)
			}
		}
	}
}

// Checks for the presence of blocks specified in the file,
// for which a name is specified, such as a group or task block.
// If the block specified in the file is in the configuration, it is ignored.
// Otherwise it will be added.
// "blockName" contains a list of available block names,
// as in the specification (group, task, etc.).
func checkNamedDublicateBlocks(
	block, file *model.TemplateBlock,
	blockName []string,
) {
	var haveBlock []string

	for _, fileBlock := range file.Block {
		if slices.Contains(blockName, fileBlock.BlockName) {

			for _, item := range block.Block {
				if item.BlockName == fileBlock.BlockName {
					if fileBlock.Name == "" {
						haveBlock = append(haveBlock, fileBlock.Name)
						continue
					}

					if item.Name == fileBlock.Name {
						haveBlock = append(haveBlock, fileBlock.Name)
					}
				}
			}
		}
	}

	for _, item := range file.Block {
		if slices.Contains(blockName, item.BlockName) {
			if !slices.Contains(haveBlock, item.Name) {
				block.Block = append(block.Block, item)
			}
		}
	}
}

// Checks for blocks in the file for which the
// specified name is not explicitly allocated from the block,
// such as a check block or a service.
// If the block specified in the file is in the configuration, it is ignored.
// Otherwise it will be added.
// The "nameKey" specifies the name of the block as a key,
// as in the specification (group, task, etc.),
// and the key name of the block (name, value, etc.).
func checkUnnamedDublicateBlocks(
	block, file *model.TemplateBlock,
	nameKey map[string]string,
) {
	var haveBlock []string

	for _, fileBlock := range file.Block {

		for blockName, name := range nameKey {
			if fileBlock.BlockName == blockName {

				for _, item := range block.Block {
					if item.BlockName == fileBlock.BlockName {
						fName := getUnnamedBlockName(name, fileBlock.Parameter)
						cName := getUnnamedBlockName(name, item.Parameter)

						if fName == "" {
							haveBlock = append(haveBlock, fileBlock.Name)
							continue
						}

						if fName == cName {
							haveBlock = append(haveBlock, fileBlock.Name)
						}
					}
				}
			}
		}
	}

	for _, item := range file.Block {
		for blockName := range nameKey {
			if item.BlockName == blockName {
				if !slices.Contains(haveBlock, item.Name) {
					block.Block = append(block.Block, item)
				}
			}
		}
	}
}

// Get the key name in a block in which
// it is not explicitly allocated from the block.
func getUnnamedBlockName(
	name string,
	parameters []map[string]interface{},
) string {
	for _, p := range parameters {
		for k, v := range p {
			if k == name {
				return v.(string)
			}
		}
	}

	return ""
}

// Checking for changes to the level block.
// Determine whether there are changes at the level of each block.
func checkFileChanges(
	block *model.TemplateBlock,
	changes *model.BlockChanges,
	blockType string,
	nameKey ...map[string]string,
) model.BlockChanges {
	var fileChanges model.TemplateBlock

	for _, fileBlock := range changes.File.Block {
		if blockType == singleType {
			if fileBlock.BlockName == block.BlockName {
				fileChanges = fileBlock
				break
			}
		}

		if blockType == namedType {
			if fileBlock.BlockName == block.BlockName {
				if fileBlock.Name == "" {
					fileChanges = fileBlock
					break
				}

				if fileBlock.Name == block.Name {
					fileChanges = fileBlock
					break
				}
			}
		}

		if blockType == unnamedType {
			if fileBlock.BlockName == block.BlockName {

				for blockName, name := range nameKey[0] {
					if fileBlock.BlockName == blockName {
						fName := getUnnamedBlockName(name, fileBlock.Parameter)
						cName := getUnnamedBlockName(name, block.Parameter)

						if fName == "" {
							fileChanges = fileBlock
							break
						}

						if fName == cName {
							fileChanges = fileBlock
							break
						}
					}
				}
			}
		}
	}

	blockChanges := model.BlockChanges{
		ProjectDirPath: changes.ProjectDirPath,
		Release:        changes.Release,
		Namespace:      changes.Namespace,
		File:           fileChanges,
		Chart:          changes.Chart,
	}

	return blockChanges
}

// Making changes to the configuration block parameters specified
// in the file and adding parameters
// if they are not specified in the configuration.
// The check is performed at the level of each block.
func setFileChanges(config, changes *model.TemplateBlock) {
	var parameters []string

	if config.BlockName != changes.BlockName || len(changes.Parameter) == 0 {
		return
	}

	for index, item := range config.Parameter {
		for configKey := range item {

			for _, p := range changes.Parameter {
				for changesKey, value := range p {
					if configKey == changesKey {
						switch value.(type) {
						case []interface{}:
							list := config.Parameter[index][configKey].([]interface{})

							for _, item := range value.([]interface{}) {
								if !slices.Contains(list, item) {
									list = append(list, item)
								}
							}

							config.Parameter[index][configKey] = list
							parameters = append(parameters, changesKey)
						default:
							config.Parameter[index][configKey] = value
							parameters = append(parameters, changesKey)
						}
					}
				}
			}
		}
	}

	for _, item := range changes.Parameter {
		for k := range item {
			if !slices.Contains(parameters, k) {
				config.Parameter = append(config.Parameter, item)
			}
		}
	}
}
