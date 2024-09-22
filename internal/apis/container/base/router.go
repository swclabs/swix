package base

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/server"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var _ = app.Router(NewRouter)

// NewRouter creates a new Base object
func NewRouter(controller IController) IRouter {
	return &Router{
		controllers: controller,
	}
}

// IRouter interface for router objects
type IRouter interface {
	server.IRouter
}

// Router struct define the Router object
type Router struct {
	controllers IController
}

// Routers implements IBase.
func (b *Router) Routers(e *echo.Echo) {
	e.GET("/docs/*any", echoSwagger.WrapHandler)

	// endpoint for common home page
	e.GET("/", Home)

	// endpoint for common routes
	e.GET("/healthcheck", b.controllers.HealthCheck)
	e.GET("/worker", b.controllers.WorkerCheck)
}
