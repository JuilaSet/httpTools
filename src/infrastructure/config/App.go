package config

type App struct {
	Port    int                `json:"port" yaml:"port"`
	Proxies []ProxyConfig      `json:"proxies" yaml:"proxies"`
	Statics []StaticFileConfig `json:"statics" yaml:"statics"`
	Uploads []UploadConfig     `json:"uploads" yaml:"uploads"`
}
