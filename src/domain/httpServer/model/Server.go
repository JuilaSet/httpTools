package model

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Server struct {
	ID        int
	Engine    *gin.Engine
	Server    *http.Server
	IsRunning bool
	*ServiceConfig
}

var n = 0
func NewHttpServer(config *ServiceConfig) *Server {
	n++
	engine := gin.New()
	config.Apply(engine)
	c := &Server{
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

func (g *Server) Quit() {
	log.Println(g.ID, "Shutdown server ...")

	// 5 秒超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		g.IsRunning = false
	}()

	if err := g.Server.Shutdown(ctx); err != nil {
		log.Fatal(g.ID, "server Shutdown error:", err)
	}
	log.Println(g.ID, "server exiting")
}

func (g *Server) Run() error {
	g.IsRunning = true

	log.Println(g.ID, "Starting server...")
	for _, v := range g.Engine.Routes() {
		fmt.Println("server:", g.ID, v.Method, v.Path)
	}
	if err := g.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	return nil
}
