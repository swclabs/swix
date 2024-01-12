package router

import (
	"github.com/gin-gonic/gin"
	"swclabs/swiftcart/delivery/http/controller"
)

func Common(e *gin.Engine) {
	r := e.Group("/v1/common")
	r.GET("/healthcheck", controller.HealthCheck)
	r.GET("/worker", controller.WorkerCheck)
	r.GET("/foo", controller.Foo)
	e.GET("/callback", controller.Auth0Callback)
}
