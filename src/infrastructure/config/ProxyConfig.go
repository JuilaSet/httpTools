package config

type ProxyConfig struct {
	Method string `json:"method" yaml:"method"`
	Target string `json:"target" yaml:"target"`
	Route  string `json:"route" yaml:"route"`
}
