package router

import (
	"swclabs/swipecore/internal/webapi/controller"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// IBase interface for router objects
type IBase interface {
	IRouter
}

// New creates a new Base object
func New(controller controller.IBase) IBase {
	return &Base{
		controllers: controller,
	}
}

var _ IBase = (*Base)(nil)

// Base struct define the Base object
type Base struct {
	controllers controller.IBase
}

// Routers implements IBase.
func (b *Base) Routers(e *echo.Echo) {
	e.GET("/docs/*any", echoSwagger.WrapHandler)

	// endpoint for common home page
	e.GET("/", controller.Home)

	// endpoint for common routes
	r := e.Group("/common")
	r.GET("/healthcheck", b.controllers.HealthCheck)
	r.GET("/worker", b.controllers.WorkerCheck)
}
