package config

type StaticFileConfig struct {
	Dir   string `json:"dir" yaml:"dir"`
	Route string `json:"route" yaml:"route"`
}
