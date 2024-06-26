package router

import (
	"swclabs/swipecore/internal/http/controller"
	"swclabs/swipecore/internal/http/middleware"

	"github.com/labstack/echo/v4"
)

const TypeAccountManagement = "AccountManagement"

type IAccountManagement interface {
	IRouter
}

type AccountManagement struct {
	controller controller.IAccountManagement
}

func NewAccountManagement(controllers controller.IAccountManagement) IAccountManagement {
	return &AccountManagement{
		controller: controllers,
	}
}

func (account *AccountManagement) Routers(e *echo.Echo) {
	user := e.Group("/users") // endpoint for users
	user.GET("", account.controller.GetMe, middleware.SessionProtected)
	user.PUT("", account.controller.UpdateUserInfo)
	user.POST("/image", account.controller.UpdateUserImage, middleware.SessionProtected)
	auth := e.Group("/auth") // endpoint for authentication
	auth.GET("", account.controller.CheckLoginEmail)
	auth.POST("/signup", account.controller.SignUp)
	auth.POST("/login", account.controller.Login)
	auth.GET("/logout", account.controller.Logout)
	auth0 := e.Group("/oauth2") // endpoint for oauth2 service
	auth0.GET("/login", controller.Auth0Login)
	e.GET("/callback", controller.Auth0Callback)
}
