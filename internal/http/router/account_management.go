package router

import (
	"swclabs/swipecore/internal/http/controller"
	"swclabs/swipecore/internal/http/middleware"

	"github.com/labstack/echo/v4"
)

const TypeAccountManagement = "AccountManagement"

type AccountManagement struct {
	controller controller.IAccountManagement
}

func NewAccountManagement(controllers *controller.AccountManagement) *AccountManagement {
	return &AccountManagement{
		controller: controllers,
	}
}

func (account *AccountManagement) Routers(e *echo.Echo) {
	user := e.Group("/users")
	user.GET("", account.controller.GetMe, middleware.SessionProtected)
	user.PUT("", account.controller.UpdateUserInfo)
	user.POST("/image", account.controller.UpdateUserImage, middleware.SessionProtected)

	auth := e.Group("/auth")
	auth.GET("", account.controller.CheckLoginEmail)
	auth.POST("/signup", account.controller.SignUp)
	auth.POST("/login", account.controller.Login)
	auth.GET("/logout", account.controller.Logout)

	auth0 := e.Group("/oauth2")
	auth0.GET("/login", controller.Auth0Login)
	e.GET("/callback", controller.Auth0Callback)
}
