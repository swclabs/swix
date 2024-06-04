package http

import (
	"os"
	"swclabs/swipecore/internal/http/router"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	_BackgroundTask(tasks ...func())
	_InitMiddleware()
	_LoggerWriter(*os.File)
	middleware(mdws ...func(*echo.Echo))
	Connect(routers router.IRouter)
	Run(string) error
}

type _Server struct {
	engine *echo.Echo
}

func NewServer(common *router.Common, docs *router.Docs) IServer {
	server := &_Server{
		engine: echo.New(),
	}
	server._InitMiddleware()
	server.Connect(common)
	server.Connect(docs)
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

func (server *_Server) Connect(routers router.IRouter) {
	routers.Routers(server.engine)
}
