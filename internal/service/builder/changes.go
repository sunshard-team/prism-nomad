package builder

import "prism/internal/model"

type ChangesStructure interface {
	Job(block model.TemplateBlock, chart model.ConfigBlock) model.TemplateBlock
	Meta(block model.TemplateBlock, chart model.ConfigBlock) model.TemplateBlock
}

type Changes struct{}

func (c *Changes) Job(
	block model.TemplateBlock,
	chart model.ConfigBlock,
) model.TemplateBlock {
Loop:
	for _, p := range chart.Parameter {
		for key, value := range p {

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
	chart model.ConfigBlock,
) model.TemplateBlock {
Loop:
	for _, p := range chart.Parameter {
		for key, value := range p {

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
