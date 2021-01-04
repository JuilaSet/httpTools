package model

import (
	"httpTools/src/domain/httpServer/model/fileUploader"
	"httpTools/src/domain/httpServer/model/proxyInfo"
	"httpTools/src/domain/httpServer/model/staticFileInfo"
	"httpTools/src/infrastructure/config"
	"strconv"
)

// builder
type Builder func(app *ServerConfig)
type Builders []Builder

func (builders Builders) Apply(model *ServerConfig) {
	for _, f := range builders {
		f(model)
	}
}

// build methods
func WithConfig(c *config.Config) Builder {
	return func(config *ServerConfig) {
		config.Port = strconv.Itoa(c.App.Port)

		// 构建代理器
		for _, p := range c.App.Proxies {
			config.Proxies = append(config.Proxies, proxyInfo.NewProxy(
				proxyInfo.WithHttpMethod(p.Method),
				proxyInfo.WithRoute(p.Route),
				proxyInfo.WithTargetURL(p.Target),
			))
		}

		// 构建静态资源
		for _, s := range c.App.Statics {
			config.Statics = append(config.Statics, staticFileInfo.NewStaticFileInfo(
				staticFileInfo.WithDir(s.Dir),
				staticFileInfo.WithRoute(s.Route),
			))
		}

		// 构建文件上传器
		for _, u := range c.App.Uploads {
			config.Uploads = append(config.Uploads, fileUploader.NewDirUploader(
				fileUploader.NewVUploadDir(u.Dir),
				fileUploader.NewVRoute(u.Route),
				fileUploader.NewVTarget(u.Target),
			))
		}
	}
}

