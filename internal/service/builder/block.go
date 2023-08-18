package builder

import (
	"fmt"
	"prism/internal/model"
)

type BlockBuilder interface {
	Job(block model.ConfigBlock, chart map[string]interface{}) model.TemplateBlock
	Group(block model.ConfigBlock) []model.TemplateBlock
}

type Block struct{}

// Get configuration block by nomad block name.
func getBlockByName(name string, job model.ConfigBlock) model.ConfigBlock {
	for _, b := range job.Block {
		if b.Name == name {
			return b
		}
	}

	return model.ConfigBlock{}
}

func (b *Block) Job(
	block model.ConfigBlock,
	chart map[string]interface{},
) model.TemplateBlock {
	job := structBuilder.Job(block)

	if len(chart) > 0 {
		for k, v := range chart {
			if k == "deploy_version" {
				i := fmt.Sprintf(`deploy_version: "%s"`, v.(string))
				job.Parameter = append(job.Parameter, i)
			}
		}
	}

	return job
}

func (b *Block) Group(block model.ConfigBlock) []model.TemplateBlock {
	blockList := make([]model.TemplateBlock, 0)

	group := getBlockByName("group", block)

	for _, item := range group.Parameter {
		for _, v := range item {
			for _, i := range v.([]interface{}) {
				group := structBuilder.Group(i)
				blockList = append(blockList, group)
			}
		}
	}

	return blockList
}
