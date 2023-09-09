package project

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Project struct{}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) CreateDefautlFile(
	embedFile embed.FS,
	embedFileName, fileName, path string,
) error {
	file, err := embedFile.Open(embedFileName)
	if err != nil {
		return fmt.Errorf("error create file %s, %s", fileName, err)
	}

	defer file.Close()

	filePath := filepath.Join(path, fileName)

	createdFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error create file %s, %s", fileName, err)
	}

	defer createdFile.Close()

	_, err = io.Copy(createdFile, file)
	if err != nil {
		return fmt.Errorf("error create file %s, %s", fileName, err)
	}

	return nil
}
