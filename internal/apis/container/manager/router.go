package manager

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/container/base"
	"swclabs/swix/internal/apis/middleware"
	"swclabs/swix/internal/apis/server"

	"github.com/labstack/echo/v4"
)

var _ = app.Router(NewRouter)

// NewManager creates a new Manager router object
func NewRouter(controllers IController) IRouter {
	return &Router{
		controller: controllers,
	}
}

// IManager interface for manager
type IRouter interface {
	server.IRouter
}

// Manager struct	implementation of IManager
type Router struct {
	controller IController
}

// Routers define route endpoint
func (r *Router) Routers(e *echo.Echo) {
	// endpoint for users
	e.GET("/users", r.controller.GetMe, middleware.SessionProtected)
	e.PUT("/users", r.controller.UpdateUserInfo)
	e.PUT("/users/image", r.controller.UpdateUserImage, middleware.SessionProtected)

	// endpoint for authentication
	e.POST("/auth", r.controller.Auth)
	e.GET("/auth/email", r.controller.CheckLoginEmail)
	e.POST("/auth/signup", r.controller.SignUp)
	e.POST("/auth/login", r.controller.Login)
	e.GET("/auth/logout", r.controller.Logout)

	// endpoint for oauth2 service
	e.GET("/callback", base.Auth0Callback)
	e.GET("/oauth2/login", base.Auth0Login)
}