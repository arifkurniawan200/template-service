package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func ReadConfig() Config {
	data, err := os.ReadFile("config/api.yaml")
	if err != nil {
		panic(err)
	}

	var configuration Config

	if err = yaml.Unmarshal(data, &configuration); err != nil {
		panic(err)
	}

	// Lowercase string values recursively
	lowerCaseConfig(&configuration)

	return configuration
}

func lowerCaseConfig(i interface{}) {
	switch v := i.(type) {
	case string:
		i = strings.ToLower(v)
	case map[interface{}]interface{}:
		for key, value := range v {
			lowerCaseConfig(key)
			lowerCaseConfig(value)
		}
	case []interface{}:
		for _, value := range v {
			lowerCaseConfig(value)
		}
	}
}
