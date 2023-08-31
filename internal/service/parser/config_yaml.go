package parser

import (
	"fmt"
	"log"
	"prism/internal/model"

	"gopkg.in/yaml.v3"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

// Parsing the chart configuration file.
func (p *Parser) ParseChart(file []byte) (map[string]interface{}, error) {
	var config map[string]interface{}

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshal file %s", err)
	}

	return config, nil
}

// Parsing the job configuration file.
func (p *Parser) ParseJob(file []byte) (model.ConfigBlock, error) {
	var config map[string]interface{}

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return model.ConfigBlock{}, fmt.Errorf(
			"error unmarshal file %s",
			err,
		)
	}

	if len(config) == 0 {
		log.Fatalf("empty job config file")
	}

	job := parseBlock("job", config["job"].(map[string]interface{}))

	return job, nil
}

// Checking if a value is a block (map[string]interface{})
// or a list of values of primitive types.
// If the value is a map[string]interface{}, then return true,
// in other cases false.
func checkBlock(value interface{}) bool {
	for _, v := range value.([]interface{}) {
		switch v.(type) {
		case map[string]interface{}:
			return true
		}
	}

	return false
}

// Parsing the configuration block.
func parseBlock(name string, config map[string]interface{}) model.ConfigBlock {
	block := model.ConfigBlock{
		Name: name,
	}

	for key, value := range config {
		parameter := make(map[string]interface{})
		parameter[key] = value

		switch v := value.(type) {
		case string, int, float32, float64, bool:
			block.Parameter = append(block.Parameter, parameter)
		case []interface{}:
			if checkBlock(value) {
				for _, item := range value.([]interface{}) {
					block.Block = append(block.Block, parseBlock(
						key,
						item.(map[string]interface{})),
					)
				}
			} else {
				block.Parameter = append(block.Parameter, parameter)
			}
		case map[string]interface{}:
			block.Block = append(block.Block, parseBlock(key, v))
		}
	}

	return block
}
