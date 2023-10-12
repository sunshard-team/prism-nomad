package deployment

import (
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"regexp"
	"strings"
)

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
	buildStructure := model.BuildStructure{
		Config:       jobConfig,
		FilesDirPath: filepath.Join(parameter.ProjectDirPath, "files"),
	}

	config = s.builder.BuildConfigStructure(buildStructure)

	//  Set changes.
	var files []model.TemplateBlock

	for _, file := range parameter.Files {
		file = filepath.Join(file)

		var (
			fileDirPath  string
			fileFullPath string
		)

		// Check the full file path or file name.
		separatorFormat, err := regexp.Compile(`\\|\/`)
		if err != nil {
			return config, fmt.Errorf(
				"failed check OS separator in file path, %s", err,
			)
		}

		findSeparator := separatorFormat.FindStringSubmatch(file)

		if len(findSeparator) > 0 {
			fileFormat, err := regexp.Compile(`([\w+-]+)\..*$`)
			if err != nil {
				return config, fmt.Errorf(
					"failed get file directory path",
				)
			}

			findFile := fileFormat.FindStringSubmatch(file)
			fileDirPath = file[:len(file)-len(findFile[0])]
			fileFullPath = file
		} else {
			fileDirPath = filepath.Join(parameter.ProjectDirPath, "files")
			fileFullPath = filepath.Join(fileDirPath, file)
		}

		// Read and parse file.
		readFile, err := os.ReadFile(fileFullPath)
		if err != nil {
			return config, fmt.Errorf("error to read job file, %s", err)
		}

		parsedFile, err := s.parser.ParseYAML(readFile)
		if err != nil {
			return config, fmt.Errorf("failed to parsing job file, %s", err)
		}

		fileConfig := s.parser.ParseConfig(
			"job",
			parsedFile["job"].(map[string]interface{}),
		)

		// Create config structure.
		buildStructure := model.BuildStructure{
			Config:       fileConfig,
			FilesDirPath: fileDirPath,
		}

		fileConfigStructure := s.builder.BuildConfigStructure(buildStructure)
		files = append(files, fileConfigStructure)
	}

	changes := model.Changes{
		Release:   parameter.Release,
		Namespace: parameter.Namespace,
		Files:     files,
		Chart:     chartConfig,
	}

	err = s.changes.SetChanges(&config, &changes)
	if err != nil {
		return config, fmt.Errorf("failed to make changes, %s", err)
	}

	return config, nil
}
