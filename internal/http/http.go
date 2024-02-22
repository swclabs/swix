package http

import (
	"github.com/labstack/echo/v4"
)

type IServer interface {
	middleware(mdws ...func(*echo.Echo))
	backgroundTask(tasks ...func())
	router(routers ...func(*echo.Echo))
	setting()
	Bootstrap(fn ...func(server IServer))
	Run(string) error
}

type _Server struct {
	engine *echo.Echo
}

func New() IServer {
	server := &_Server{
		engine: echo.New(),
	}
	server.setting()
	return server
}

func (server *_Server) middleware(mdws ...func(*echo.Echo)) {
	for _, m := range mdws {
		m(server.engine)
	}
}

func (server *_Server) backgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (server *_Server) router(routers ...func(*echo.Echo)) {
	for _, r := range routers {
		r(server.engine)
	}
}
