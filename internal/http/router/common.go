package router

import (
	"github.com/labstack/echo/v4"
	"github.com/swclabs/swipe-api/internal/http/controller"
)

func Common(e *echo.Echo) {
	r := e.Group("common")
	r.GET("/healthcheck", controller.HealthCheck)
	r.GET("/worker", controller.WorkerCheck)
	r.GET("/foo", controller.Foo)
}
