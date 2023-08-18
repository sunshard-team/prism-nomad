package output

import (
	"bytes"
	"fmt"
	"os"
	"prism/internal/model"
	"prism/internal/templates"
	"text/template"

	"github.com/Masterminds/sprig"
)

type Output struct{}

func NewOutput() *Output {
	return &Output{}
}

func (s *Output) CreateNomadConfigFile(
	name, path string,
	config model.TemplateBlock,
) error {
	// Add dynamic indentation.
	// Add "include" function, to replace "template".

	tmpl := template.New("config").Funcs(sprig.FuncMap())
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
		return fmt.Errorf("error create include function for templates %s", err)
	}

	// Create nomad config template.

	jobTmpl, err := createTemplate(tmpl)
	if err != nil {
		return err
	}

	groupTmpl, err := createTemplate(jobTmpl)
	if err != nil {
		return err
	}

	taskGroup, err := createTemplate(groupTmpl)
	if err != nil {
		return err
	}

	// Create new .nomad.hcl file.
	filePath := fmt.Sprintf("%s/%s.nomad.hcl", path, name)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error create nomad configuration file, %s", err)
	}

	defer file.Close()

	// Write data into nomad configuration template file.
	err = taskGroup.ExecuteTemplate(file, "block", config)
	if err != nil {
		return fmt.Errorf("error execute nomad configuration template, %v", err)
	}

	return nil
}

func createTemplate(tmpl *template.Template) (*template.Template, error) {
	t, err := template.Must(tmpl.Clone()).ParseFS(
		templates.TemplateFile,
		"nomad_block_config.tmpl",
	)

	if err != nil {
		return nil, fmt.Errorf("error create template %s", err)
	}

	return t, nil
}
