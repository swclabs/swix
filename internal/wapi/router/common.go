// Package router This file contains the common router for the wapi service.
package router

import (
	"swclabs/swipecore/internal/wapi/controller"

	"github.com/labstack/echo/v4"
)

// ICommon extends IRouter interface
type ICommon interface {
	IRouter
}

// Common router implementation ICommon
type Common struct {
	controllers controller.ICommon
}

// NewCommon creates a new Common router object
func NewCommon(controllers controller.ICommon) ICommon {
	return &Common{
		controllers: controllers,
	}
}

// Routers define route endpoint
func (c *Common) Routers(e *echo.Echo) {
	r := e.Group("/common")
	r.GET("/healthcheck", c.controllers.HealthCheck)
	r.GET("/worker", c.controllers.WorkerCheck)
	r.GET("/foo", controller.Foo)
}
