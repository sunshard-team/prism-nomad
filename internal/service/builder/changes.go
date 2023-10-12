package builder

import (
	"prism/internal/model"
	"slices"
)

var (
	single  = "single"
	unnamed = "unnamed"
	named   = "named"
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
		Release:   changes.Release,
		Namespace: changes.Namespace,
		File:      model.TemplateBlock{},
		Topic:     changes.Topic,
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
func checkSingleBlocks(block, file *model.TemplateBlock, blockType []string) {
	var haveBlock []string

	for _, fileBlock := range file.Block {
		if slices.Contains(blockType, fileBlock.Type) {
			for _, item := range block.Block {
				if item.Type == fileBlock.Type {
					haveBlock = append(haveBlock, fileBlock.Type)
				}
			}
		}
	}

	for _, item := range file.Block {
		if slices.Contains(blockType, item.Type) {
			if !slices.Contains(haveBlock, item.Type) {
				block.Block = append(block.Block, item)
			}
		}
	}
}

// Checks for the presence of blocks specified in the file,
// for which a name is specified, such as a group or task block.
// If the block specified in the file is in the configuration, it is ignored.
// Otherwise it will be added.
func checkNamedDublicateBlocks(
	block, file *model.TemplateBlock,
	blockType []string,
) {
	var haveBlock []string

	for _, fileBlock := range file.Block {
		if slices.Contains(blockType, fileBlock.Type) {

			for _, item := range block.Block {
				if item.Type == fileBlock.Type {
					if fileBlock.Label == "" {
						haveBlock = append(haveBlock, fileBlock.Label)
						continue
					}

					if item.Label == fileBlock.Label {
						haveBlock = append(haveBlock, fileBlock.Label)
					}
				}
			}
		}
	}

	for _, item := range file.Block {
		if slices.Contains(blockType, item.Type) {
			if !slices.Contains(haveBlock, item.Label) {
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
// The "nameKey" specifies the type of the block as a key,
// as in the specification (group, task, etc.),
// and the key name of the block (name, value, etc.).
func checkUnnamedDublicateBlocks(
	block, file *model.TemplateBlock,
	nameKey map[string]string,
) {
	var haveBlock []string

	for _, fileBlock := range file.Block {

		for blockType, key := range nameKey {
			if fileBlock.Type == blockType {

				for _, item := range block.Block {
					if item.Type == fileBlock.Type {
						fileKey := getUnnamedBlockName(key, fileBlock.Parameter)
						blockKey := getUnnamedBlockName(key, item.Parameter)

						if fileKey == "" {
							haveBlock = append(haveBlock, fileBlock.Label)
							continue
						}

						if fileKey == blockKey {
							haveBlock = append(haveBlock, fileBlock.Label)
						}
					}
				}
			}
		}
	}

	for _, item := range file.Block {
		for blockType := range nameKey {
			if item.Type == blockType {
				if !slices.Contains(haveBlock, item.Label) {
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
	blockKind string,
	nameKey ...map[string]string,
) model.BlockChanges {
	var fileChanges model.TemplateBlock

	for _, fileBlock := range changes.File.Block {
		if blockKind == single {
			if fileBlock.Type == block.Type {
				fileChanges = fileBlock
				break
			}
		}

		if blockKind == named {
			if fileBlock.Type == block.Type {
				if fileBlock.Label == "" {
					fileChanges = fileBlock
					break
				}

				if fileBlock.Label == block.Label {
					fileChanges = fileBlock
					break
				}
			}
		}

		if blockKind == unnamed {
			if fileBlock.Type == block.Type {

				for blockType, label := range nameKey[0] {
					if fileBlock.Type == blockType {
						fileLabel := getUnnamedBlockName(label, fileBlock.Parameter)
						blockLabel := getUnnamedBlockName(label, block.Parameter)

						if fileLabel == "" {
							fileChanges = fileBlock
							break
						}

						if fileLabel == blockLabel {
							fileChanges = fileBlock
							break
						}
					}
				}
			}
		}
	}

	blockChanges := model.BlockChanges{
		Release:   changes.Release,
		Namespace: changes.Namespace,
		File:      fileChanges,
		Topic:     changes.Topic,
	}

	return blockChanges
}

// Making changes to the configuration block parameters specified
// in the file and adding parameters
// if they are not specified in the configuration.
// The check is performed at the level of each block.
func setFileChanges(config, changes *model.TemplateBlock) {
	var parameters []string

	if config.Type != changes.Type || len(changes.Parameter) == 0 {
		return
	}

	for index, item := range config.Parameter {
		for configKey := range item {

			for _, p := range changes.Parameter {
				for key, value := range p {
					if configKey == key {
						switch value.(type) {
						case []interface{}:
							list := config.Parameter[index][configKey].([]interface{})

							for _, item := range value.([]interface{}) {
								if !slices.Contains(list, item) {
									list = append(list, item)
								}
							}

							config.Parameter[index][configKey] = list
							parameters = append(parameters, key)
						default:
							config.Parameter[index][configKey] = value
							parameters = append(parameters, key)
						}
					}
				}
			}
		}
	}

	for _, parameter := range changes.Parameter {
		for key := range parameter {
			if !slices.Contains(parameters, key) {
				config.Parameter = append(config.Parameter, parameter)
			}
		}
	}
}
