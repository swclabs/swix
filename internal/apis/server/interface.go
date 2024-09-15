// Package server interface
package server

import (
	"os"
	"swclabs/swix/internal/apis/router"

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
