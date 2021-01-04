package test

import (
	"httpTools/src/domain/httpServer"
	"httpTools/src/infrastructure"
	"httpTools/src/infrastructure/config"
	"httpTools/src/infrastructure/event"
	"testing"
)

func TestHttpUploaderServer(t *testing.T) {
	// 配置项目
	configFilename := "config.yml"
	appConfig := config.NewAppConfig(configFilename)

	// 事件存储
	emitter := event.NewEmitter()
	emitter.Start()

	// 构建服务
	httpEvent := httpServer.NewEvent(emitter, appConfig)

	// 唤起服务
	httpEvent.Emit()

	infrastructure.OSWait()
}

