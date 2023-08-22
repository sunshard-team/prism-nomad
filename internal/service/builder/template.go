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

	// Group.
	groups := blockBuilder.Group(config)

	job.Block = append(job.Block, groups...)

	return job
}
