package router

import (
	"github.com/gin-gonic/gin"
	controller2 "swclabs/swiftcart/delivery/http/controller"
	"swclabs/swiftcart/delivery/http/middleware"
)

func Users(e *gin.Engine) {
	user := e.Group("/v1/users")
	user.GET("/", middleware.SessionProtected, controller2.GetMe)
	user.PUT("/", controller2.UpdateUserInfo)
	user.POST("/image", middleware.SessionProtected, controller2.UpdateUserImage)
}

func Auth(e *gin.Engine) {
	auth := e.Group("/v1/auth")
	auth.POST("/signup", controller2.SignUp)
	auth.POST("/login", controller2.Login)
	auth.GET("/logout", controller2.Logout)
}

func OAuth2(e *gin.Engine) {
	auth0 := e.Group("/v1/oauth2")
	auth0.GET("/login", controller2.Auth0Login)
}
