// Package router implements the router interface
package router

import (
	"github.com/labstack/echo/v4"
)

// IRouter interface for router objects
type IRouter interface {
	Routers(e *echo.Echo)
}
