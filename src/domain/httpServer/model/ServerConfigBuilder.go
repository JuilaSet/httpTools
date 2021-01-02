package model

import (
	"httpTools/src/domain/httpServer/model/proxyInfo"
	"httpTools/src/domain/httpServer/model/staticFileInfo"
	"httpTools/src/infrastructure/config"
	"strconv"
)

// builder
type Builder func(app *ServiceConfig)
type Builders []Builder

func (builders Builders) Apply(model *ServiceConfig) {
	for _, f := range builders {
		f(model)
	}
}

// build methods
func WithConfig(c *config.Root) Builder {
	return func(config *ServiceConfig) {
		config.Port = strconv.Itoa(c.Config.Port)

		// 构建代理器
		for _, p := range c.Config.Proxies {
			config.Proxies = append(config.Proxies, proxyInfo.NewProxy(
				proxyInfo.WithHttpMethod(p.Method),
				proxyInfo.WithRoute(p.Route),
				proxyInfo.WithTargetURL(p.Target),
			))
		}

		// 构建静态资源
		for _, s := range c.Config.Statics {
			config.Statics = append(config.Statics, staticFileInfo.NewStaticFileInfo(
				staticFileInfo.WithDir(s.Dir),
				staticFileInfo.WithRoute(s.Route),
			))
		}
	}
}

