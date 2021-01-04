package model

import (
	"github.com/gin-gonic/gin"
	"httpTools/src/domain/httpServer/model/fileUploader"
	"httpTools/src/domain/httpServer/model/proxyInfo"
	"httpTools/src/domain/httpServer/model/staticFileInfo"
	"httpTools/src/domain/httpServer/vo"
	"log"
)

// aggregate root
type ServerConfig struct {
	*vo.VServerInfo `json:"service_info"`
	Proxies         proxyInfo.Proxies       `json:"proxies"`
	Statics         staticFileInfo.Statics  `json:"statics"`
	Uploads         fileUploader.DirUploads `json:"uploads"`
}

func NewServiceConfig(builder ...Builder) *ServerConfig {
	c := &ServerConfig{
		VServerInfo: vo.NewVServiceInfo("80"),
		Proxies: make(proxyInfo.Proxies, 0),
		Statics: make(staticFileInfo.Statics, 0),
		Uploads: make(fileUploader.DirUploads, 0),
	}
	Builders(builder).Apply(c)
	return c
}

func (info *ServerConfig) Apply(engine *gin.Engine) {
	log.Println("Apply Service App ...")
	info.Proxies.Apply(engine)
	info.Statics.Apply(engine)
	info.Uploads.Apply(engine)
}
