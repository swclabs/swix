package router

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/controller"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var _ = app.Router(New)

// New creates a new Base object
func New(controller controller.IBaseController) IBaseRouter {
	return &BaseRouter{
		controllers: controller,
	}
}

// IBaseRouter interface for router objects
type IBaseRouter interface {
	IRouter
}

// BaseRouter struct define the BaseRouter object
type BaseRouter struct {
	controllers controller.IBaseController
}

// Routers implements IBase.
func (b *BaseRouter) Routers(e *echo.Echo) {
	e.GET("/docs/*any", echoSwagger.WrapHandler)

	// endpoint for common home page
	e.GET("/", controller.Home)

	// endpoint for common routes
	r := e.Group("/common")
	r.GET("/healthcheck", b.controllers.HealthCheck)
	r.GET("/worker", b.controllers.WorkerCheck)
}
