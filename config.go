package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Config Application Config
type Config struct {
	API struct {
		URL      string `yaml:"url"`
		Key      string `yaml:"key"`
		Username string `yaml:"username"`
	}
}

// NewConfig - Create a new app config
func NewConfig(configFile string) (cfg *Config, err error) {
	var config *Config
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(byteValue, &config); err != nil {
		return nil, err
	}
	return config, nil
}
