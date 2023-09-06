package builder

import "prism/internal/model"

type ChangesStructure interface {
	Job(block model.TemplateBlock, chart map[string]interface{}) model.TemplateBlock
	Meta(block model.TemplateBlock, chart map[string]interface{}) model.TemplateBlock
}

type Changes struct{}

func (c *Changes) Job(
	block model.TemplateBlock,
	chart map[string]interface{},
) model.TemplateBlock {
	if len(chart) > 0 {
	Loop:
		for key, value := range chart {
			if key == "type" {
				for i := range block.Parameter {
					for k := range block.Parameter[i] {
						if k == key {
							block.Parameter[i][k] = value
							break Loop
						}
					}
				}
			}
		}
	}

	return block
}

func (c *Changes) Meta(
	block model.TemplateBlock,
	chart map[string]interface{},
) model.TemplateBlock {
	if len(chart) > 0 {
	Loop:
		for key, value := range chart {
			if key == "deploy_version" {
				for i := range block.Parameter {
					for k := range block.Parameter[i] {
						if k == "run_uuid" {
							block.Parameter[i][k] = value
							break Loop
						}
					}
				}
			}
		}
	}

	return block
}
