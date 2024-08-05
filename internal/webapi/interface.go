// Package webapi interface
package webapi

import (
	"os"
	"swclabs/swix/internal/types"
	"swclabs/swix/internal/webapi/router"

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
