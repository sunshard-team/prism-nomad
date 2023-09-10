package service

import (
	"embed"
	"prism/internal/model"
	"prism/internal/service/builder"
	"prism/internal/service/deployment"
	"prism/internal/service/output"
	"prism/internal/service/parser"
	"prism/internal/service/project"
)

type Project interface {
	CreateDefautlFile(
		embedFile embed.FS,
		embedFileName, fileName, path string,
	) error
}

type Output interface {
	// Returns the formated job configuration of the nomad.
	OutputConfig(config model.TemplateBlock) (string, error)

	// Creates a nomad configuration file in .nomad.hcl format.
	CreateConfigFile(name, path string, config model.TemplateBlock) error
}

type Parser interface {
	// Parsing the chart configuration file.
	ParseChart(file []byte) (model.ConfigBlock, error)

	// Parsing the job configuration file.
	ParseJob(file []byte) (model.ConfigBlock, error)
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
) *Service {
	return &Service{
		Project:    project.NewProject(),
		Output:     output.NewOutput(),
		Parser:     parser.NewParser(),
		Builder:    builder.NewStructureBuilder(),
		Deployment: deployment.NewDeployment(*p, *b),
	}
}
