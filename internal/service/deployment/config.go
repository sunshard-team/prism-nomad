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
		return config, fmt.Errorf("error to read chart file, %s", err)
	}

	parsedChartConfig, err := s.parser.ParseYAML(chartFile)
	if err != nil {
		if err != nil {
			return config, fmt.Errorf("failed to parsing chart file, %s", err)
		}
	}

	chartConfig := s.parser.ParseConfig("chart", parsedChartConfig)

	// Job.
	jobFileName := "config.yaml"
	jobPath := filepath.Join(parameter.ProjectDirPath, jobFileName)

	jobFile, err := os.ReadFile(jobPath)
	if err != nil {
		return config, fmt.Errorf("error to read job file, %s", err)
	}

	parsedJobConfig, err := s.parser.ParseYAML(jobFile)
	if err != nil {
		return config, fmt.Errorf("failed to parsing job file, %s", err)
	}

	jobConfig := s.parser.ParseConfig(
		"job",
		parsedJobConfig["job"].(map[string]interface{}),
	)

	// Config structure.
	config = s.builder.BuildConfigStructure(
		jobConfig, chartConfig, parameter.ProjectDirPath,
	)

	return config, nil
}
