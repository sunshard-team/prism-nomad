// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package deployment

import (
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Returns the configuration structure.
func (s *Deployment) CreateConfigStructure(
	parameter model.ConfigParameter,
) ([]model.TemplateBlock, error) {
	var configList []model.TemplateBlock

	// Pack.
	packFileName := "pack.yaml"
	packPath := filepath.Join(parameter.ProjectDirPath, packFileName)

	packFile, err := os.ReadFile(packPath)
	if err != nil {
		return configList, fmt.Errorf("error to read pack file, %s", err)
	}

	var packConfig = &model.Pack{}

	err = yaml.Unmarshal([]byte(packFile), packConfig)
	if err != nil {
		return configList, fmt.Errorf("failed to parsing pack file, %s", err)
	}

	// Job config.
	configFileName := "config.yaml"
	configPath := filepath.Join(parameter.ProjectDirPath, configFileName)
	filesPath := filepath.Join(parameter.ProjectDirPath, "files")

	configStructure, err := s.BuildConfigStructure(configPath, "job")
	if err != nil {
		return configList, err
	}

	// Set changes.
	config, err := s.SetChanges(filesPath, parameter, packConfig, configStructure)
	if err != nil {
		return configList, err
	}

	// Create dependencies configuration structure.
	if len(packConfig.Dependencies) > 0 {
		for _, dependencyJob := range packConfig.Dependencies {
			configFileName := "config.yaml"
			configPath := filepath.Join(dependencyJob.Path, configFileName)
			filesPath := filepath.Join(dependencyJob.Path, "files")

			configStructure, err := s.BuildConfigStructure(configPath, "job")
			if err != nil {
				return configList, err
			}

			parameter.Files = dependencyJob.Files

			config, err := s.SetChanges(filesPath, parameter, packConfig, configStructure)
			if err != nil {
				return configList, err
			}

			configList = append(configList, config)
		}
	}

	configList = append(configList, config)
	return configList, nil
}

func (s *Deployment) SetChanges(
	filesDirPath string,
	parameter model.ConfigParameter,
	packConfig *model.Pack,
	config model.TemplateBlock,
) (model.TemplateBlock, error) {
	// Parsing files.
	var files []model.TemplateBlock

	for _, file := range parameter.Files {
		file = filepath.Join(file)

		_, fileFullPath, err := s.CheckFileName(file, parameter.ProjectDirPath)
		if err != nil {
			return config, fmt.Errorf("could not verify file name, %s", err)
		}

		fileConfigStructure, err := s.BuildConfigStructure(fileFullPath, "job")
		if err != nil {
			return config, err
		}

		files = append(files, fileConfigStructure)
	}

	// Set changes.
	changes := model.Changes{
		Release:      parameter.Release,
		Namespace:    parameter.Namespace,
		Files:        files,
		FilesDirPath: filesDirPath,
		Pack:         *packConfig,
		EnvFilePath:  parameter.EnvFilePath,
		EnvVars:      parameter.EnvVars,
	}

	err := s.changes.SetChanges(&config, &changes)
	if err != nil {
		return config, fmt.Errorf("failed to make changes, %s", err)
	}

	return config, nil
}

// Parsing the configuration file and creating a structured job configuration.
func (s *Deployment) BuildConfigStructure(path, blockType string) (model.TemplateBlock, error) {
	var config model.TemplateBlock

	content, err := s.ParseFile(path)
	if err != nil {
		return config, fmt.Errorf("parse error, %s", err)
	}

	parsedConfig := s.parser.ParseConfig(
		blockType,
		content["job"].(map[string]interface{}),
	)

	buildStructure := model.BuildStructure{
		Config: parsedConfig,
	}

	config = s.builder.BuildConfigStructure(buildStructure)
	return config, nil
}

// Read and parse file.
// Returns the file contents, hierarchically sorted into blocks.
func (s *Deployment) ParseFile(fileFullPath string) (map[string]interface{}, error) {
	var parsedContent map[string]interface{}

	content, err := os.ReadFile(fileFullPath)
	if err != nil {
		return parsedContent, err
	}

	parsedContent, err = s.parser.ParseYAML(content)
	if err != nil {
		return parsedContent, fmt.Errorf("failed to parsing file %s, %s", fileFullPath, err)
	}

	return parsedContent, nil
}

// Check the full file path or file name.
func (s *Deployment) CheckFileName(
	file, projectDirPath string,
) (fileDirPath, fileFullPath string, err error) {
	separatorFormat, err := regexp.Compile(`\\|\/`)
	if err != nil {
		return fileDirPath, fileFullPath, fmt.Errorf(
			"failed check OS separator in file path, %s", err,
		)
	}

	findSeparator := separatorFormat.FindStringSubmatch(file)

	if len(findSeparator) > 0 {
		fileFormat, err := regexp.Compile(`([\w+-]+)\..*$`)
		if err != nil {
			return fileDirPath, fileFullPath, fmt.Errorf(
				"failed get file directory path",
			)
		}

		findFile := fileFormat.FindStringSubmatch(file)
		fileDirPath = file[:len(file)-len(findFile[0])]
		fileFullPath = file
	} else {
		fileDirPath = filepath.Join(projectDirPath, "files")
		fileFullPath = filepath.Join(fileDirPath, file)
	}

	return fileDirPath, fileFullPath, nil
}
