package config

type Config struct {
	Port    int                `json:"port" yaml:"port"`
	Proxies []ProxyConfig      `json:"proxies" yaml:"proxies"`
	Statics []StaticFileConfig `json:"statics" yaml:"statics"`
}
