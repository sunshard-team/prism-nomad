package deployment

import (
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"prism/internal/service/builder"
	"prism/internal/service/parser"
	"strings"
)

type Deployment struct {
	parser  parser.Parser
	builder builder.StructureBuilder
}

func NewDeployment(
	parser parser.Parser,
	builder builder.StructureBuilder,
) *Deployment {
	return &Deployment{
		parser:  parser,
		builder: builder,
	}
}

// Returns the configuration structure.
func (s *Deployment) CreateConfigStructure(
	parameter model.ConfigParameter,
) (model.TemplateBlock, error) {
	var config model.TemplateBlock

	// Chart.
	chartName := strings.ReplaceAll(parameter.ProjectDir, "-", "_")
	chartFileName := fmt.Sprintf("%s.yaml", chartName)
	chartPath := filepath.Join(parameter.ProjectDirPath, chartFileName)

	chartFile, err := os.ReadFile(chartPath)
	if err != nil {
		return config, fmt.Errorf("failed read chart file")
	}

	chartConfig, err := s.parser.ParseChart(chartFile)
	if err != nil {
		return config, fmt.Errorf("failed parse chart config file, %s", err)
	}

	// Job.
	jobFileName := "config.yaml"
	jobPath := filepath.Join(parameter.ProjectDirPath, jobFileName)

	jobFile, err := os.ReadFile(jobPath)
	if err != nil {
		return config, fmt.Errorf("failed read job file")
	}

	jobConfig, err := s.parser.ParseJob(jobFile)
	if err != nil {
		return config, fmt.Errorf("filed parse job config file, %s", err)
	}

	// Config structure.
	config = s.builder.BuildConfigStructure(
		jobConfig, chartConfig, parameter.ProjectDirPath,
	)

	return config, nil
}
