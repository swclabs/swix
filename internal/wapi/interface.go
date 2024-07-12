// Package wapi interface
package wapi

import (
	"os"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/wapi/router"

	"github.com/labstack/echo/v4"
)

// IServer interface represents all server method
type IServer interface {
	backgroundTask(tasks ...func())
	initMiddleware()
	loggerWriter(*os.File)
	middleware(mdws ...func(*echo.Echo))
	Connect(routers router.IRouter)
	Routes() []string
	Run(string) error
}

// NewBaseAdapter returns a new adapter wrapping around the given server
func NewBaseAdapter(server IServer) types.IAdapter {
	adapter := &_Adapter{
		server: server,
	}
	return adapter
}
