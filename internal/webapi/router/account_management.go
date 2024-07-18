// Package router account management implementation
// File account_management.go defines routes for APIs related to account management,
// such as login, signup, logout, and user information.
package router

import (
	"swclabs/swipecore/internal/webapi/controller"
	"swclabs/swipecore/internal/webapi/middleware"

	"github.com/labstack/echo/v4"
)

// IAccountManagement interface for account management
type IAccountManagement interface {
	IRouter
}

// AccountManagement struct	implementation of IAccountManagement
type AccountManagement struct {
	controller controller.IAccountManagement
}

// NewAccountManagement creates a new AccountManagement router object
func NewAccountManagement(controllers controller.IAccountManagement) IAccountManagement {
	return &AccountManagement{
		controller: controllers,
	}
}

// Routers define route endpoint
func (account *AccountManagement) Routers(e *echo.Echo) {
	user := e.Group("/users") // endpoint for users
	user.GET("", account.controller.GetMe, middleware.SessionProtected)
	user.PUT("", account.controller.UpdateUserInfo)
	user.PUT("/image", account.controller.UpdateUserImage, middleware.SessionProtected)
	auth := e.Group("/auth") // endpoint for authentication
	auth.GET("", account.controller.CheckLoginEmail)
	auth.POST("/signup", account.controller.SignUp)
	auth.POST("/login", account.controller.Login)
	auth.GET("/logout", account.controller.Logout)
	auth0 := e.Group("/oauth2") // endpoint for oauth2 service
	auth0.GET("/login", controller.Auth0Login)
	e.GET("/callback", controller.Auth0Callback)
}
