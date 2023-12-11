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

func setEnvVar(config *model.TemplateBlock, filePath string, env map[string]string) error {
	envFormat, err := regexp.Compile(`\${(PRISM_([\w+-]+))}|\${((PRISM_([\w+-]+))\|default=([\w+-]+))}`)
	if err != nil {
		return fmt.Errorf("error set env, %s", err)
	}

	envDefaultFormat, err := regexp.Compile(`\|default=([\w+-]+)`)
	if err != nil {
		return fmt.Errorf("error set env, %s", err)
	}

	if config.Label != "" {
		newLabel, err := findEnv(config.Label, filePath, envFormat, envDefaultFormat)
		if err != nil {
			return fmt.Errorf("failed set environment variable, %s", err)
		}

		config.Label = newLabel
	}

	if len(config.Parameter) > 0 {
		for index, item := range config.Parameter {
			for key, value := range item {

				switch v := value.(type) {
				case string:
					newValue, err := findEnv(v, filePath, envFormat, envDefaultFormat)
					if err != nil {
						return fmt.Errorf("failed set environment variable, %s", err)
					}

					config.Parameter[index][key] = envTyping(newValue)
				case []interface{}:
					var paramenters []interface{}

					for _, item := range v {
						newValue, err := findEnv(item.(string), filePath, envFormat, envDefaultFormat)
						if err != nil {
							return fmt.Errorf("failed set environment variable, %s", err)
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
			err := setEnvVar(&config.Block[index], "", map[string]string{})
			if err != nil {
				return fmt.Errorf("error set env, %s", err)
			}
		}
	}

	return nil
}

func findEnv(origin, path string, format, defFormat *regexp.Regexp) (string, error) {
	envGroup := format.FindAllStringSubmatch(origin, -1)

	for _, item := range envGroup {
		emptyDefaultValue := true

		// Finding a variable by key "PRISM_".
		// Поиск переменной по ключу "PRISM_".
		value, err := getEnv(item[1], path)

		if value == nil {
			value, err = getEnv(item[4], path)
		}

		if err != nil {
			return "", err
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
				emptyDefaultValue = false
			}
		}

		// If the variable is not found, add it to the list of not found to display an error.
		// Если переменная не найдена, добавьте ее в список не найденных, чтобы вывести ошибку.
		if value == nil && emptyDefaultValue {
			mEnv := fmt.Sprint(item[1], item[4])

			if !slices.Contains(missingENV, mEnv) {
				missingENV = append(missingENV, mEnv)
			}
		}

	}

	return origin, nil
}

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

func getEnv(name, filePath string) (interface{}, error) {
	vp := viper.New()

	vp.SetConfigFile(filePath)
	vp.AutomaticEnv()

	err := vp.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error get env %s, %s", name, err)
	}

	env := vp.Get(name)
	return env, nil
}
