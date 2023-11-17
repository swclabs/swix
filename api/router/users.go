package router

import (
	"example/swiftcart/api/controller"
	"example/swiftcart/api/middleware"

	"github.com/gin-gonic/gin"
)

func Users(e *gin.Engine) {
	r := e.Group("/v1/auth")
	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)

	usr := e.Group("/v1")
	usr.GET("/users", middleware.Protected, controller.GetMe)
}
