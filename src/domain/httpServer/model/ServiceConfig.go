package model

import (
	"github.com/gin-gonic/gin"
	"httpTools/src/domain/httpServer/model/proxyInfo"
	"httpTools/src/domain/httpServer/model/staticFileInfo"
	"httpTools/src/domain/httpServer/vo"
	"log"
)

// aggregate root
type ServiceConfig struct {
	*vo.VServerInfo `json:"service_info"`
	Proxies         proxyInfo.Proxies      `json:"proxies"`
	Statics         staticFileInfo.Statics `json:"statics" yaml:"statics"`
}

func NewServiceConfig(builder ...Builder) *ServiceConfig {
	c := &ServiceConfig{
		VServerInfo: vo.NewVServiceInfo(""),
		Proxies:     proxyInfo.Proxies{},
	}
	Builders(builder).Apply(c)
	return c
}

func (info *ServiceConfig) Apply(engine *gin.Engine) {
	log.Println("Apply Service Config ...")
	info.Proxies.Apply(engine)
	info.Statics.Apply(engine)
}
