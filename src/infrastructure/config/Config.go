package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Root struct {
	Config Config `json:"app" yaml:"app"`
}

func NewAppConfig(filepath string) *Root {
	c := &Root{}
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
