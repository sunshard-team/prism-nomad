package service

import (
	"prism/internal/model"
	"prism/internal/service/builder"
	"prism/internal/service/deployment"
	"prism/internal/service/output"
	"prism/internal/service/parser"
	"prism/internal/service/project"
)

type Project interface {
	// Creates a project.
	Create(name string) (string, error)
}

type Output interface {
	// Returns the formated job configuration of the nomad.
	OutputConfig(config model.TemplateBlock) (string, error)

	// Creates a nomad configuration file in .nomad.hcl format.
	CreateConfigFile(name, path string, config model.TemplateBlock) error
}

type Parser interface {
	// Parsing the YAML configuration file.
	ParseYAML(file []byte) (map[string]interface{}, error)

	// Parsing the configuration map. Assembles a block structure.
	ParseConfig(name string, config map[string]interface{}) model.ConfigBlock
}

type Builder interface {
	// Builds and returns a job configuration structure.
	BuildConfigStructure(
		job, chart model.ConfigBlock,
		projectDirPath string,
	) model.TemplateBlock
}

type Deployment interface {
	// Returns the configuration structure.
	CreateConfigStructure(
		parameter model.ConfigParameter,
	) (model.TemplateBlock, error)
}

type Service struct {
	Project    Project
	Output     Output
	Parser     Parser
	Builder    Builder
	Deployment Deployment
}

func NewService(
	p *parser.Parser,
	b *builder.StructureBuilder,
	o *output.Output,
) *Service {
	return &Service{
		Project:    project.NewProject(*p, *b, *o),
		Output:     output.NewOutput(),
		Parser:     parser.NewParser(),
		Builder:    builder.NewStructureBuilder(),
		Deployment: deployment.NewDeployment(*p, *b),
	}
}
