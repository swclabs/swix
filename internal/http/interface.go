package http

import (
	"github.com/labstack/echo/v4"
	"os"
	"swclabs/swipecore/internal/http/router"
)

// IAdapter interface, used to connect to server instance
type IAdapter interface {
	Run(addr string) error
	Routers() []string
}

type IServer interface {
	backgroundTask(tasks ...func())
	initMiddleware()
	loggerWriter(*os.File)
	middleware(mdws ...func(*echo.Echo))
	Routes() []string
	Connect(routers router.IRouter)
	Run(string) error
}

func NewBaseAdapter(server IServer) IAdapter {
	adapter := &_Adapter{
		server: server,
	}
	return adapter
}
