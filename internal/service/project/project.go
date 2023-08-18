package project

import (
	"embed"
	"fmt"
	"io"
	"os"
)

type Project struct{}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) CreateDefautlFile(embedFile embed.FS, fileName, path string) error {
	file, err := embedFile.Open(fileName)
	if err != nil {
		return fmt.Errorf("error create file %s, %s", fileName, err)
	}

	defer file.Close()

	filePath := fmt.Sprintf("%s/%s", path, fileName)

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
