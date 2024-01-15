package http

import (
	"swclabs/swiftcart/internal/config"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	middleware(...func(*gin.Engine))
	backgroundTask(...func())
	router(...func(*gin.Engine))
	Run(string) error
}

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	if config.StageStatus != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Server{
		engine: gin.Default(),
	}
}

func (server *Server) middleware(mdws ...func(*gin.Engine)) {
	for _, m := range mdws {
		m(server.engine)
	}
}

func (server *Server) backgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (server *Server) router(routers ...func(*gin.Engine)) {
	for _, r := range routers {
		r(server.engine)
	}
}
