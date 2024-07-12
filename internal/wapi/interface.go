package wapi

import (
	"os"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/wapi/router"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	backgroundTask(tasks ...func())
	initMiddleware()
	loggerWriter(*os.File)
	middleware(mdws ...func(*echo.Echo))
	Connect(routers router.IRouter)
	Routes() []string
	Run(string) error
}

func NewBaseAdapter(server IServer) types.IAdapter {
	adapter := &_Adapter{
		server: server,
	}
	return adapter
}
