package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"prism/internal/model"
	"prism/internal/templates"
	"text/template"

	"gopkg.in/yaml.v3"
)

type OutputService struct{}

func NewOutputService() *OutputService {
	return &OutputService{}
}

func (s *OutputService) CreateNomadConfiguration(
	name, path, from string,
) (bool, error) {
	var created bool

	// Read config.yaml file.
	data, err := os.ReadFile(from)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Parse config.yaml into a Config structure.
	var config model.Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// Create new .nomad file.
	filePath := fmt.Sprintf("%s/%s", path, name)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create %s: %v", name, err)
	}

	defer file.Close()

	// Create a nomad template from the config structure.
	tmpl := template.Must(
		template.New("nomad_template.tmpl").Funcs(template.FuncMap{
			"marshalArgs":    marshalArgs,
			"marshalVolumes": marshalVolumes,
		}).ParseFS(templates.TemplateFile, "nomad_template.tmpl"),
	)

	err = tmpl.Execute(file, config)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}

	fmt.Printf("%s generated successfully!\n", name)
	return created, nil
}

func marshalArgs(args []string) (string, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func marshalVolumes(volumes []string) (string, error) {
	b, err := json.Marshal(volumes)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
