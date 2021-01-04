package main

import (
	"httpTools/src/domain/configFileNotifier"
	"httpTools/src/domain/httpServer"
	"httpTools/src/infrastructure"
	"httpTools/src/infrastructure/config"
	"httpTools/src/infrastructure/event"
)

func main() {
	// 配置项目
	configFilename := "config.yml"
	appConfig := config.NewAppConfig(configFilename)

	// 事件存储
	emitter := event.NewEmitter()
	emitter.Start()

	// 构建服务
	httpEvent := httpServer.NewEvent(emitter, appConfig)
	fileNotifierEvent := configFileNotifier.NewEvent(emitter, configFilename)

	// 唤起服务
	httpEvent.Emit()
	fileNotifierEvent.Emit()

	infrastructure.OSWait()
}
