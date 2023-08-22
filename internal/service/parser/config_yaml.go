package parser

import (
	"fmt"
	"prism/internal/model"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseChart(file []byte) (map[string]interface{}, error) {
	var config map[string]interface{}

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshal file %s", err)
	}

	return config, nil
}

func (p *Parser) ParseJobConfig(file []byte) (model.ConfigBlock, error) {
	var config map[string]interface{}

	job := model.ConfigBlock{
		Name: "job",
	}

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return job, fmt.Errorf("error unmarshal file %s", err)
	}

	for k, v := range config["job"].(map[string]interface{}) {
		i := make(map[string]interface{})
		i[k] = v

		switch v.(type) {
		case string:
			job.Parameter = append(job.Parameter, i)
		case int:
			job.Parameter = append(job.Parameter, i)
		case float32:
			job.Parameter = append(job.Parameter, i)
		case float64:
			job.Parameter = append(job.Parameter, i)
		case bool:
			job.Parameter = append(job.Parameter, i)
		case []interface{}:
			check := checkBlock(v)
			if check {
				job.Block = append(job.Block, parseBlock(k, i))
				continue
			}

			job.Parameter = append(job.Parameter, i)
		case map[string]interface{}:
			job.Block = append(job.Block, parseBlock(k, i))
		default:
			continue
		}
	}

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

func parseBlock(name string, blockMap map[string]interface{}) model.ConfigBlock {
	block := model.ConfigBlock{
		Name: name,
	}

	for _, value := range blockMap {
		rt := reflect.TypeOf(value)

		switch rt.Kind() {
		case reflect.String:
			block.Parameter = append(block.Parameter, blockMap)
		case reflect.Int:
			block.Parameter = append(block.Parameter, blockMap)
		case reflect.Float32:
			block.Parameter = append(block.Parameter, blockMap)
		case reflect.Float64:
			block.Parameter = append(block.Parameter, blockMap)
		case reflect.Bool:
			block.Parameter = append(block.Parameter, blockMap)
		case reflect.Slice:
			block.Parameter = append(block.Parameter, blockMap)
		case reflect.Map:
			for k, v := range value.(map[string]interface{}) {
				i := make(map[string]interface{})
				i[k] = v

				block.Block = append(block.Block, parseBlock(k, i))
			}
		default:
			continue
		}
	}

	return block
}
