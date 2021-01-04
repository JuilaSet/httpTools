package config

type UploadConfig struct {
	Dir    string `json:"dir" yaml:"dir"`
	Target string `json:"target" yaml:"target"`
	Route  string `json:"route" yaml:"route"`
}
