package ConfigFileReload

import (
	"errors"
	"httpTools/src/infrastructure/config"
	"httpTools/src/infrastructure/event/vo"
)

type ConfigFileReload struct {
	data *config.Config
}

func NewConfigFileReload(data *config.Config) *ConfigFileReload {
	return &ConfigFileReload{data: data}
}

func FromData(data vo.VData) (c *ConfigFileReload) {
	if data == nil {
		panic(errors.New("data is nil"))
	}
	switch data.(type) {
	case *config.Config:
		return NewConfigFileReload(data.(*config.Config))
	default:
		panic(errors.New("data type not support yet"))
	}
}

func (c *ConfigFileReload) Data() *config.Config {
	return c.data
}
