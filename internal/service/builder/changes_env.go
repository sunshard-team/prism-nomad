// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package builder

import (
	"fmt"
	"prism/internal/model"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

var missingEnvVars []string

// Searching for environment variables in the configuration structure
// to replace them with values from one of
// the sources in which the variable is specified.
func findAndReplaceEnvVars(
	config *model.TemplateBlock,
	filePath string,
	envVars map[string]string,
) error {
	err := setEnvVar(config, filePath, envVars)
	if err != nil {
		return fmt.Errorf(
			"an error occurred while searching and inserting environment variables, %s",
			err,
		)
	}

	if len(missingEnvVars) > 0 {
		var missingVars string

		for i, m := range missingEnvVars {
			if len(missingEnvVars)-1 != i {
				missingVars += fmt.Sprintf("%s, ", m)
			} else {
				missingVars += m
			}
		}

		return fmt.Errorf("environment variables not found: %v", missingVars)
	}

	return nil
}

// Searching for environment variables with the "PRISM_" key
// in the configuration file and replacing them with the found values
// from local environment variables, a file with variables
// or values specified in the deployment command flag.
func setEnvVar(
	config *model.TemplateBlock,
	filePath string,
	envVars map[string]string,
) error {
	// Find format.
	envFormat, err := regexp.Compile(
		`\${(PRISM_([\w+-]+))}|\${((PRISM_([\w+-]+))\|default=([\w+-]+))}`,
	)

	if err != nil {
		return fmt.Errorf("error parse regex, %s", err)
	}

	envDefaultFormat, err := regexp.Compile(`\|default=([\w+-]+)`)
	if err != nil {
		return fmt.Errorf("error parse regex, %s", err)
	}

	// Finding and replace a variable in a block label.
	if config.Label != "" {
		newLabel, err := replaceEnvVar(
			config.Label,
			filePath,
			envVars,
			envFormat,
			envDefaultFormat,
		)

		if err != nil {
			return fmt.Errorf(
				"failed set environment variable for label in block %s, %s",
				config.Type, err,
			)
		}

		config.Label = newLabel
	}

	// Finding and replace a variables in a parameters.
	if len(config.Parameter) > 0 {
		for index, item := range config.Parameter {
			for key, value := range item {

				switch v := value.(type) {
				case string:
					newValue, err := replaceEnvVar(
						v,
						filePath,
						envVars,
						envFormat,
						envDefaultFormat,
					)

					if err != nil {
						return fmt.Errorf(
							"failed set environment variable for paramenter %s in block %s, %s",
							key, config.Type, err,
						)
					}

					config.Parameter[index][key] = envTyping(newValue)
				case []interface{}:
					var paramenters []interface{}

					for _, item := range v {
						newValue, err := replaceEnvVar(
							item.(string),
							filePath,
							envVars,
							envFormat,
							envDefaultFormat,
						)

						if err != nil {
							return fmt.Errorf(
								"failed set environment variable for paramenter %s in block %s, %s",
								key, config.Type, err,
							)
						}

						paramenters = append(paramenters, envTyping(newValue))
					}

					config.Parameter[index][key] = paramenters
				}
			}
		}
	}

	if len(config.Block) > 0 {
		for index := range config.Block {
			err := setEnvVar(&config.Block[index], filePath, envVars)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Searches for an environment variable with the "PRISM_" key
// and replace it with the value of a variable found in the local environment,
// a file with variables, or specified in the deployment command flag.
func replaceEnvVar(
	origin, path string,
	envVars map[string]string,
	format, defFormat *regexp.Regexp,
) (string, error) {
	envGroup := format.FindAllStringSubmatch(origin, -1)

	for _, item := range envGroup {
		missingDefaultValue := true

		// Finding a variable by key "PRISM_".
		value, err := getEnv(item[1], path)
		if err != nil {
			return "", err
		}

		// Finding a variable by key "PRISM_", with default value.
		if value == nil {
			value, err = getEnv(item[4], path)
			if err != nil {
				return "", err
			}
		}

		// Set the env variable from the set flags.
		if len(envVars) > 0 {
			for k, v := range envVars {
				if k == fmt.Sprint(item[1], item[4]) {
					value = v
				}
			}
		}

		// Replacing variables with values.
		if value != nil {
			// Environment variable.
			origin = strings.Replace(origin, item[0], fmt.Sprint(value), -1)
		} else {
			// Default value.
			envDefaultGroup := defFormat.FindAllStringSubmatch(item[0], -1)

			for _, defaultItem := range envDefaultGroup {
				origin = strings.Replace(origin, item[0], defaultItem[1], -1)
				missingDefaultValue = false
			}
		}

		// If the variable is not found, add it to the list of not found to display an error.
		if value == nil && missingDefaultValue {
			mEnv := fmt.Sprint(item[1], item[4])

			if !slices.Contains(missingEnvVars, mEnv) {
				missingEnvVars = append(missingEnvVars, mEnv)
			}
		}
	}

	return origin, nil
}

// Checking the value type of a variable.
// Returns string, int, float32, float64, bool.
func envTyping(value string) any {
	var a any

	if v, err := strconv.Atoi(value); err == nil {
		a = v
	} else if v, err := strconv.ParseBool(value); err == nil {
		a = v
	} else if v, err := strconv.ParseFloat(value, 32); err == nil {
		a = v
	} else if v, err := strconv.ParseFloat(value, 64); err == nil {
		a = v
	} else {
		return value
	}

	return a
}

// Search for an environment variable in a file or local environment.
// Priority is given to the variables specified in the file,
// i.e. if a variable is specified both in the local environment
// and in a file at the same time,
// the value will be taken from the variable specified in the file.
func getEnv(name, filePath string) (interface{}, error) {
	var envValue interface{}
	vp := viper.New()

	// Get environment variable from file.
	if filePath != "" {
		vp.SetConfigFile(filePath)

		err := vp.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("error get env %s, %s", name, err)
		}

		envValue = vp.Get(name)

		if envValue != nil {
			return envValue, nil
		}
	}

	// If the environment variable is not in the file,
	// try to find it in the local environment.
	vp.AutomaticEnv()
	envValue = vp.Get(name)

	return envValue, nil
}
