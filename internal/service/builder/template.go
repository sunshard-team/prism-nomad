package builder

import (
	"prism/internal/model"
)

type TemplateBuilder struct{}

func NewTemplateBuilder() *TemplateBuilder {
	return &TemplateBuilder{}
}

func (b *TemplateBuilder) BuildConfigTemplate(
	jobConfig model.ConfigBlock,
	chartConfig map[string]interface{},
) model.TemplateBlock {
	config := jobConfig
	chart := chartConfig

	// Job.
	job := blockBuilder.Job(config, chart)
	var jobBlock []model.TemplateBlock

	// Group.
	groups := blockBuilder.Group(config)
	// var groupBlock []model.TemplateBlock

	jobBlock = append(jobBlock, groups...)
	job.Block = append(job.Block, jobBlock...)

	return job
}
