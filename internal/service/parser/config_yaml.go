// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"fmt"
	"prism/internal/model"

	"gopkg.in/yaml.v3"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

// Parsing the YAML configuration file.
func (p *Parser) ParseYAML(file []byte) (map[string]interface{}, error) {
	var config map[string]interface{}

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf("parsing file error, %s", err)
	}

	if len(config) == 0 {
		return config, fmt.Errorf("file is empty")
	}

	return config, nil
}

// Parsing the configuration map.
// Assembles a block structure.
func (p *Parser) ParseConfig(
	blockType string,
	config map[string]interface{},
) model.ConfigBlock {
	block := model.ConfigBlock{
		Type: blockType,
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
					block.Block = append(block.Block, p.ParseConfig(
						key,
						item.(map[string]interface{})),
					)
				}
			} else {
				block.Parameter = append(block.Parameter, parameter)
			}
		case map[string]interface{}:
			block.Block = append(block.Block, p.ParseConfig(key, v))
		}
	}

	return block
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
