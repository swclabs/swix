package router

import (
	"example/komposervice/api/controller"

	"github.com/gin-gonic/gin"
)

func Auth(e *gin.Engine) {
	r := e.Group("/v1/auth")
	r.POST("/sign-up", controller.SignUp)
	r.POST("/sign-in", controller.SignIn)
}
