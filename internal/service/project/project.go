package project

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"prism/config"
	"prism/internal/service/builder"
	"prism/internal/service/output"
	"prism/internal/service/parser"
	"strings"
)

type Project struct {
	parser  parser.Parser
	builder builder.StructureBuilder
	output  output.Output
}

func NewProject(
	parser parser.Parser,
	builder builder.StructureBuilder,
	output output.Output,
) *Project {
	return &Project{
		parser:  parser,
		builder: builder,
		output:  output,
	}
}

// Creates a new project.
func (s *Project) Create(name string) (string, error) {
	// Get root dir path.
	rootDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed project initialization, %s", err)
	}

	// Create project directories.
	projectName := "prism"

	if name != "" {
		projectName = name
	}

	projectDirPath := filepath.Join(rootDir, projectName)

	projectFileDir := "files"
	fileDirPath := filepath.Join(projectDirPath, projectFileDir)

	chartName := strings.ReplaceAll(projectName, "-", "_")
	chartFileName := fmt.Sprintf("%s.yaml", chartName)

	configName := "config"
	configFileName := fmt.Sprintf("%s.yaml", configName)

	dirStat, err := os.Stat(projectDirPath)
	if err != nil || !dirStat.IsDir() {
		err = os.MkdirAll(
			filepath.Join(projectName, projectFileDir),
			0700,
		)

		if err != nil {
			return "", fmt.Errorf("failed project initialization, %s", err)
		}
	}

	// Create default project files.
	defaultFile := map[string]string{
		"prism.yaml":  chartFileName,
		"config.yaml": configFileName,
	}

	for k, v := range defaultFile {
		err = createFile(config.ConfigFile, k, v, projectDirPath)

		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
	}

	// additional files.
	err = createFile(config.ConfigFile,
		"load_balancer.conf",
		"load_balancer.conf",
		fileDirPath,
	)

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	// Create .nomad.hcl configuration file.
	// Parse chart config file.
	chartFilePath := filepath.Join(projectDirPath, chartFileName)
	chartFile, err := os.ReadFile(chartFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read chart file, %s", err)
	}

	parsedChartConfig, err := s.parser.ParseYAML(chartFile)
	if err != nil {
		return "", fmt.Errorf("failed to parsing chart file, %s", err)
	}

	chartConfig := s.parser.ParseConfig("chart", parsedChartConfig)

	// Parse job config file.
	jobFilePath := filepath.Join(projectDirPath, configFileName)
	jobFile, err := os.ReadFile(jobFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read job file, %s", err)
	}

	parsedJobConfig, err := s.parser.ParseYAML(jobFile)
	if err != nil {
		return "", fmt.Errorf("failed to parsing job file, %s", err)
	}

	jobConfig := s.parser.ParseConfig(
		"job",
		parsedJobConfig["job"].(map[string]interface{}),
	)

	configStructure := s.builder.BuildConfigStructure(
		jobConfig,
		chartConfig,
		projectDirPath,
	)

	err = s.output.CreateConfigFile(
		configName,
		projectDirPath,
		configStructure,
	)

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return projectName, nil
}

// Creates project file.
func createFile(
	embedFile embed.FS,
	embedFileName, fileName, path string,
) error {
	file, err := embedFile.Open(embedFileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s, %s", fileName, err)
	}

	defer file.Close()

	filePath := filepath.Join(path, fileName)

	createdFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s, %s", fileName, err)
	}

	defer createdFile.Close()

	_, err = io.Copy(createdFile, file)
	if err != nil {
		return fmt.Errorf("failed to create file %s, %s", fileName, err)
	}

	return nil
}
