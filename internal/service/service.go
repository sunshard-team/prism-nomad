package service

import (
	"embed"
	"prism/internal/model"
	"prism/internal/service/builder"
	"prism/internal/service/output"
	"prism/internal/service/parser"
	"prism/internal/service/project"
)

type Project interface {
	CreateDefautlFile(embedFile embed.FS, fileName, path string) error
}

type Output interface {
	CreateNomadConfigFile(name, path string, config model.TemplateBlock) error
}

type Parser interface {
	ParseChart(file []byte) (map[string]interface{}, error)
	ParseJob(file []byte) (model.ConfigBlock, error)
}

type Builder interface {
	BuildConfigTemplate(
		jobConfig model.ConfigBlock,
		chartConfig map[string]interface{},
		projectPath string,
	) model.TemplateBlock
}

type Service struct {
	Project Project
	Output  Output
	Parser  Parser
	Builder Builder
}

func NewService() *Service {
	return &Service{
		Project: project.NewProject(),
		Output:  output.NewOutput(),
		Parser:  parser.NewParser(),
		Builder: builder.NewTemplateBuilder(),
	}
}
