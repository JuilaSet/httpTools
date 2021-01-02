package main

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"httpProxyDDD/src/domain"
	"httpProxyDDD/src/domain/fileNotifier"
	"httpProxyDDD/src/domain/fileNotifier/model/fileWatcher"
	"httpProxyDDD/src/domain/httpServer"
	"httpProxyDDD/src/domain/httpServer/model"
	"httpProxyDDD/src/infrastructure/config"
	"httpProxyDDD/src/infrastructure/event"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		r := gin.Default()
		r.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "test",
			})
		})
		r.GET("/test/api", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "api",
			})
		})
		r.Run("127.0.0.1:8080")
	}()

	// 消息队列
	emitter := event.NewEmitter()
	emitter.StartListen()

	var server domain.IServer
	emitter.On("RestartServer", func(data event.Data) {
		if server != nil {
			server.Quit()
		}

		// 读取配置
		app := model.NewApp(model.WithConfig(config.NewAppConfig("config.yml")))
		server = httpServer.NewHttpServer(app)

		s, _ := json.MarshalIndent(app, "", "\t")
		log.Println("app info", string(s))

		// 启动服务器
		go func() {
			if err := server.Run(); err != nil {
				panic(err)
			}
		}()
	})

	var fileWatchService domain.IServer
	emitter.On("StartWatcher", func(data event.Data) {
		if fileWatchService != nil {
			fileWatchService.Quit()
		}

		// 监视器服务
		fileWatchService = fileNotifier.NewFileWatcherService(fileWatcher.NewFileStatusWatcher(
			fileWatcher.WithFilename("config.yml"),
			fileWatcher.WithWriteHandlers(func(event fsnotify.Event) {
				log.Println("写入文件 : ", event.Name)
				emitter.Emit("RestartServer", event)
			}),
		))

		// 启动监视器
		go func() {
			if err := fileWatchService.Run(); err != nil {
				log.Fatal(err)
			}
		}()
	})

	// 唤起服务
	emitter.Emit("RestartServer", true)
	emitter.Emit("StartWatcher", true)

	// 等待系统取消
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
