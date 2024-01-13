package router

import (
	"swclabs/swiftcart/internal/http/controller"
	"swclabs/swiftcart/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func Users(e *gin.Engine) {
	user := e.Group("/v1/users")
	user.GET("/", middleware.SessionProtected, controller.GetMe)
	user.PUT("/", controller.UpdateUserInfo)
	user.POST("/image", middleware.SessionProtected, controller.UpdateUserImage)
}

func Auth(e *gin.Engine) {
	auth := e.Group("/v1/auth")
	auth.POST("/signup", controller.SignUp)
	auth.POST("/login", controller.Login)
	auth.GET("/logout", controller.Logout)
}

func OAuth2(e *gin.Engine) {
	auth0 := e.Group("/v1/oauth2")
	auth0.GET("/login", controller.Auth0Login)
}
