package http

import (
	"swclabs/swiftcart/internal/config"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	middleware(mdws ...func(*gin.Engine))
	backgroundTask(tasks ...func())
	router(routers ...func(*gin.Engine))
	prepare()
	InitAccountManagement()
	Run(string) error
}

type _Server struct {
	engine *gin.Engine
}

func New() IServer {
	if config.StageStatus != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := &_Server{
		engine: gin.Default(),
	}
	server.prepare()
	return server
}

func (server *_Server) middleware(mdws ...func(*gin.Engine)) {
	for _, m := range mdws {
		m(server.engine)
	}
}

func (server *_Server) backgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (server *_Server) router(routers ...func(*gin.Engine)) {
	for _, r := range routers {
		r(server.engine)
	}
}
