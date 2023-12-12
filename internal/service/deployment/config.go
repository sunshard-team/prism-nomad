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
)

// Returns the configuration structure.
func (s *Deployment) CreateConfigStructure(
	parameter model.ConfigParameter,
) (model.TemplateBlock, error) {
	var config model.TemplateBlock

	// Pack.
	packFileName := "pack.yaml"
	packPath := filepath.Join(parameter.ProjectDirPath, packFileName)

	packFile, err := os.ReadFile(packPath)
	if err != nil {
		return config, fmt.Errorf("error to read pack file, %s", err)
	}

	parsedPackConfig, err := s.parser.ParseYAML(packFile)
	if err != nil {
		if err != nil {
			return config, fmt.Errorf("failed to parsing pack file, %s", err)
		}
	}

	packConfig := s.parser.ParseConfig("pack", parsedPackConfig)

	// Job config.
	configFileName := "config.yaml"
	configPath := filepath.Join(parameter.ProjectDirPath, configFileName)

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("error to read job file, %s", err)
	}

	parsedConfig, err := s.parser.ParseYAML(configFile)
	if err != nil {
		return config, fmt.Errorf("failed to parsing job config file, %s", err)
	}

	jobConfig := s.parser.ParseConfig(
		"job",
		parsedConfig["job"].(map[string]interface{}),
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
		Release:     parameter.Release,
		Namespace:   parameter.Namespace,
		Files:       files,
		Pack:        packConfig,
		EnvFilePath: parameter.EnvFilePath,
		EnvVars:     parameter.EnvVars,
	}

	err = s.changes.SetChanges(&config, &changes)
	if err != nil {
		return config, fmt.Errorf("failed to make changes, %s", err)
	}

	return config, nil
}
