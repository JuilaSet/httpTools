package infrastructure

import (
	"os"
	"os/signal"
)

func OSWait() {
	// 等待系统取消
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
