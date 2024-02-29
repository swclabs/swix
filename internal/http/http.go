package http

import (
	"os"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	middleware(mdws ...func(*echo.Echo))
	router(routers ...func(*echo.Echo))
	_BackgroundTask(tasks ...func())
	_InitMiddleware()
	_LoggerWriter(*os.File)
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
	server._InitMiddleware()
	server.Bootstrap(CommonModule)
	return server
}

func (server *_Server) middleware(mdws ...func(*echo.Echo)) {
	for _, m := range mdws {
		m(server.engine)
	}
}

func (server *_Server) _BackgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (server *_Server) router(routers ...func(*echo.Echo)) {
	for _, r := range routers {
		r(server.engine)
	}
}
