// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package project

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"prism/config"
	"strings"
)

type Project struct{}

func NewProject() *Project {
	return &Project{}
}

// Creates a new project.
func (s *Project) Create(name string) (string, error) {
	// Get root dir path.
	rootDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get path, %s", err)
	}

	// Create project directories.
	projectName := "prism"

	if name != "" {
		projectName = name
	}

	projectDirPath := filepath.Join(rootDir, projectName)

	projectFileDir := "files"
	fileDirPath := filepath.Join(projectDirPath, projectFileDir)

	packFileName := "pack.yaml"
	configFileName := "config.yaml"

	dirStat, err := os.Stat(projectDirPath)
	if err != nil || !dirStat.IsDir() {
		err = os.MkdirAll(
			filepath.Join(projectName, projectFileDir),
			0700,
		)

		if err != nil {
			return "", fmt.Errorf("failed to check project directory, %s", err)
		}
	}

	// Create default project files.
	defaultFile := map[string]string{
		"pack.yaml":   packFileName,
		"config.yaml": configFileName,
	}

	for k, v := range defaultFile {
		err = createFile(config.ConfigFile, k, v, projectDirPath)

		if err != nil {
			return "", err
		}
	}

	err = createFile(config.ConfigFile,
		"load_balancer.conf",
		"load_balancer.conf",
		fileDirPath,
	)

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	// Set pack name.
	packPath := filepath.Join(projectDirPath, packFileName)
	err = setPackName(projectName, packPath)
	if err != nil {
		return "", err
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

func setPackName(name, filePath string) error {
	packFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error to read pack file, %s", err)
	}

	lines := strings.Split(string(packFile), "\n")

	for i, l := range lines {
		if strings.Contains(l, "name: \"PRISM_PACK_NAME\"") {
			lines[i] = fmt.Sprintf("name: \"%s\"", name)
		}
	}

	output := strings.Join(lines, "\n")

	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("error writing data to pack file, %s", err)
	}

	return nil
}
