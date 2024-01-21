package router

import (
	"github.com/swclabs/swipe-api/internal/http/controller"
	"github.com/swclabs/swipe-api/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type AccountManagement struct {
	controller controller.IAccountManagement
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		controller: controller.NewAccountManagement(),
	}
}

func (account *AccountManagement) Users(e *gin.Engine) {
	user := e.Group("/v1/users")
	user.GET("/", middleware.SessionProtected, account.controller.GetMe)
	user.PUT("/", account.controller.UpdateUserInfo)
	user.POST("/image", middleware.SessionProtected, account.controller.UpdateUserImage)
}

func (account *AccountManagement) Auth(e *gin.Engine) {
	auth := e.Group("/v1/auth")
	auth.GET("/", account.controller.CheckLoginEmail)
	auth.POST("/signup", account.controller.SignUp)
	auth.POST("/login", account.controller.Login)
	auth.GET("/logout", account.controller.Logout)
}

func (account *AccountManagement) OAuth2(e *gin.Engine) {
	auth0 := e.Group("/v1/oauth2")
	auth0.GET("/login", controller.Auth0Login)
	e.GET("/callback", controller.Auth0Callback)
}
