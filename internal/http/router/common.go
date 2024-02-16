package router

import (
	"github.com/swclabs/swipe-api/internal/http/controller"

	"github.com/gin-gonic/gin"
)

func Common(e *gin.Engine) {
	r := e.Group("common")
	r.GET("/healthcheck", controller.HealthCheck)
	r.GET("/worker", controller.WorkerCheck)
	r.GET("/foo", controller.Foo)
}
