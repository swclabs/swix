package router

import (
	"swclabs/swiftcart/api/controller"
	"swclabs/swiftcart/api/middleware"

	"github.com/gin-gonic/gin"
)

func Users(e *gin.Engine) {
	auth := e.Group("/v1/auth")
	auth.POST("/signup", controller.SignUp)
	auth.POST("/login", controller.Login)
	auth.GET("/logout", controller.Logout)

	auth0 := e.Group("/v1/auth0")
	auth0.GET("/login", controller.Auth0Login)

	user := e.Group("/v1")
	user.GET("/users", middleware.SessionProtected, controller.GetMe)
	user.PUT("/users", controller.UpdateUserInfo)

}
