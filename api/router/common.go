package router

import (
	"example/swiftcart/api/controller"

	"github.com/gin-gonic/gin"
)

func Common(e *gin.Engine) {
	r := e.Group("/v1/common")
	r.GET("/healthcheck", controller.HealthCheck)
	r.GET("/worker", controller.WorkerCheck)
	r.GET("/foo", controller.Foo)
}
