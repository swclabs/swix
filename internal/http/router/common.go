package router

import (
	"swclabs/swipecore/internal/http/controller"

	"github.com/labstack/echo/v4"
)

const TypeCommon = "common"

type Common struct {
	controllers controller.ICommon
}

func NewCommon(controllers controller.ICommon) IRouter {
	return &Common{
		controllers: controllers,
	}
}

func (c *Common) Routers(e *echo.Echo) {
	r := e.Group("/common")
	r.GET("/healthcheck", c.controllers.HealthCheck)
	r.GET("/worker", c.controllers.WorkerCheck)
	r.GET("/foo", controller.Foo)
}
