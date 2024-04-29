package http

import (
	"os"
	"swclabs/swipecore/internal/http/router"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	middleware(mdws ...func(*echo.Echo))
	connect(routers router.IRouter)
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

func (server *_Server) connect(routers router.IRouter) {
	routers.Routers(server.engine)
}
