package http

import (
	"os"
	"swclabs/swipecore/internal/http/router"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	middleware(mdws ...func(*echo.Echo))
	_BackgroundTask(tasks ...func())
	_InitMiddleware()
	_LoggerWriter(*os.File)
	Connect(routers router.IRouter)
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
	server.Connect(router.New(router.TypeCommon))
	server.Connect(router.New(router.TypeDocs))
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
