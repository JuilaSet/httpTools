package main

import (
	"httpTools/src/domain/fileNotifier"
	"httpTools/src/domain/httpServer"
	"httpTools/src/infrastructure"
	"httpTools/src/infrastructure/event"
)

func main() {
	// 事件存储
	emitter := event.NewEmitter()
	emitter.Start()

	httpServer.NewEvent(emitter)
	fileNotifier.NewEvent(emitter)

	// 唤起服务
	emitter.Emit(httpServer.EventName, true)
	emitter.Emit(fileNotifier.EventName, true)

	infrastructure.OSWait()
}
