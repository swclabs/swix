package router

import (
	"github.com/labstack/echo/v4"
	"swclabs/swipe-api/internal/http/controller"
	"swclabs/swipe-api/internal/http/middleware"
)

type AccountManagement struct {
	controller controller.IAccountManagement
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		controller: controller.NewAccountManagement(),
	}
}

func (account *AccountManagement) Users(e *echo.Echo) {
	user := e.Group("/users")
	user.GET("", account.controller.GetMe, middleware.SessionProtected)
	user.PUT("", account.controller.UpdateUserInfo)
	user.POST("/image", account.controller.UpdateUserImage, middleware.SessionProtected)
}

func (account *AccountManagement) Auth(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.GET("", account.controller.CheckLoginEmail)
	auth.POST("/signup", account.controller.SignUp)
	auth.POST("/login", account.controller.Login)
	auth.GET("/logout", account.controller.Logout)
}

func (account *AccountManagement) OAuth2(e *echo.Echo) {
	auth0 := e.Group("/oauth2")
	auth0.GET("/login", controller.Auth0Login)
	e.GET("/callback", controller.Auth0Callback)
}
