// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package output

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"prism/internal/templates"
	"text/template"

	"github.com/Masterminds/sprig"
)

type Output struct{}

func NewOutput() *Output {
	return &Output{}
}

// Returns the formated job configuration of the nomad.
func (s *Output) OutputConfig(config model.TemplateBlock) (string, error) {
	var buf bytes.Buffer

	configTemplate, err := createTemplate()
	if err != nil {
		return "", fmt.Errorf(
			"error write template to buffer, %s", err,
		)
	}

	// Write data to buffer.
	err = configTemplate.ExecuteTemplate(&buf, "block", config)
	if err != nil {
		return "", fmt.Errorf(
			"error write template to buffer, %s", err,
		)
	}

	return buf.String(), nil
}

// Creates a nomad configuration file in .nomad.hcl format.
func (s *Output) CreateConfigFile(
	name, path string,
	config model.TemplateBlock,
) error {
	configTemplate, err := createTemplate()
	if err != nil {
		return fmt.Errorf("error create nomad configuration file, %s", err)
	}

	// Create new .nomad.hcl file.
	fileName := fmt.Sprintf("%s.nomad.hcl", name)
	filePath := filepath.Join(path, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error create nomad configuration file, %s", err)
	}

	defer file.Close()

	// Write data into nomad configuration template file.
	err = configTemplate.ExecuteTemplate(file, "block", config)
	if err != nil {
		return fmt.Errorf("error execute nomad configuration template, %v", err)
	}

	return nil
}

// Creates a nomad configuration template.
func createTemplate() (*template.Template, error) {
	// Add dynamic indentation.
	// Add "include" function, to replace "template".
	tmpl := template.New("config").Funcs(sprig.FuncMap())
	tmpl.Funcs(template.FuncMap{
		"getValue": getValue,
	})

	var funcMap template.FuncMap = map[string]interface{}{}

	funcMap["include"] = func(name string, data interface{}) (string, error) {
		buf := bytes.NewBuffer(nil)
		if err := tmpl.ExecuteTemplate(buf, name, data); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	tmpl, err := tmpl.Funcs(sprig.TxtFuncMap()).Funcs(funcMap).ParseFS(
		templates.TemplateFile,
		"nomad_block_config.tmpl",
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error when creating include function for templates %s", err,
		)
	}

	// Create nomad configuration template.
	configTemplate, err := template.Must(tmpl.Clone()).ParseFS(
		templates.TemplateFile,
		"nomad_block_config.tmpl",
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error when creating nomad configuration template %s", err,
		)
	}

	return configTemplate, nil
}

func getValue(block string, value map[string]interface{}) string {
	var parameter string

	for k, v := range value {
		switch v := v.(type) {
		case string:
			if block == "template" {
				if k == "data" {
					parameter = fmt.Sprintf("%s = <<EOH\n%v\nEOH", k, v)
				} else {
					parameter = fmt.Sprintf(`%s = "%v"`, k, v)
				}
			} else {
				parameter = fmt.Sprintf(`%s = "%v"`, k, v)
			}
		case int:
			parameter = fmt.Sprintf("%s = %v", k, v)
		case bool:
			parameter = fmt.Sprintf("%s = %v", k, v)
		case []interface{}:
			listValue := make([]string, 0)

			for index, item := range v {
				switch item := item.(type) {
				case string:
					if index+1 == len(v) {
						listValue = append(listValue, fmt.Sprintf(`"%s"`, item))
					} else {
						listValue = append(listValue, fmt.Sprintf(`"%s",`, item))
					}
				case int:
					if index+1 == len(v) {
						listValue = append(listValue, fmt.Sprintf("%v", item))
					} else {
						listValue = append(listValue, fmt.Sprintf("%v,", item))
					}
				}
			}

			parameter = fmt.Sprintf("%s = %v", k, listValue)
		}
	}

	return parameter
}
