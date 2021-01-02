package httpServer

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"httpProxyDDD/src/domain/httpServer/model"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	ID        int
	Engine    *gin.Engine
	Server    *http.Server
	IsRunning bool
	*model.ServiceConfig
}

var n = 0
func NewHttpServer(config *model.ServiceConfig) *HttpServer {
	n++
	engine := gin.New()
	config.Apply(engine)
	c := &HttpServer{
		ID:            n,
		ServiceConfig: config,
		Engine:        engine,
		Server: &http.Server{
			Addr:    ":" + config.Port,
			Handler: engine,
		},
	}
	return c
}

func (g *HttpServer) Quit() {
	log.Println(g.ID, "Shutdown Server ...")

	// 5 秒超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		g.IsRunning = false
	}()

	if err := g.Server.Shutdown(ctx); err != nil {
		log.Fatal(g.ID, "Server Shutdown error:", err)
	}
	log.Println(g.ID, "Server exiting")
}

func (g *HttpServer) Run() error {
	g.IsRunning = true

	log.Println(g.ID, "Starting Server...")
	for _, v := range g.Engine.Routes() {
		fmt.Println("server:", g.ID, v.Method, v.Path)
	}
	if err := g.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	return nil
}
