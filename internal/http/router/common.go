package router

import (
	"swclabs/swipecore/internal/http/controller"

	"github.com/labstack/echo/v4"
)

const TypeCommon = "common"

type Common struct{}

func newCommon() *Common {
	return &Common{}
}

func NewCommon() *Common {
	return &Common{}
}

func (c *Common) Routers(e *echo.Echo) {
	r := e.Group("/common")
	r.GET("/healthcheck", controller.HealthCheck)
	r.GET("/worker", controller.WorkerCheck)
	r.GET("/foo", controller.Foo)
}
