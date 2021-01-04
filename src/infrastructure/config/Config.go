package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	App App `json:"app" yaml:"app"`
}

func NewAppConfig(filepath string) *Config {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}
	return c
}
