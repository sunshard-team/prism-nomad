package builder

import (
	"prism/internal/model"
)

type StructureBuilder interface {
	Artifact(block []map[string]interface{}) model.TemplateBlock
	Job(block model.ConfigBlock) model.TemplateBlock
	Group(block map[string]interface{}) model.TemplateBlock
}

type Structure struct{}

func (s *Structure) Artifact(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)
	var internalBlock []model.TemplateBlock

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "destination", "mode", "source":
				parameters = append(parameters, i)
			case "options":
				options := s.artifactBlock(k, i)
				internalBlock = append(internalBlock, options)
			case "headers":
				headers := s.artifactBlock(k, i)
				internalBlock = append(internalBlock, headers)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "artifact",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) artifactBlock(
	name string,
	block map[string]interface{},
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		if k == name {
			parameters = append(parameters, i)
			break
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Job(block model.ConfigBlock) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"all_at_once",
		"datacenters",
		"node_pool",
		"namespace",
		"priority",
		"region",
		"type",
		"vault_token",
		"consul_token",
	}

	for _, item := range block.Parameter {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			if k == "name" {
				name = v.(string)
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, i)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "job",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Group(block map[string]interface{}) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"count",
		"shutdown_delay",
		"stop_after_client_disconnect",
		"max_client_disconnect",
	}

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		if k == "name" {
			name = v.(string)
		}

		for _, p := range parameterName {
			switch k {
			case p:
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "group",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}
