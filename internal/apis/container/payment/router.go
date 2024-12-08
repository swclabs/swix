package payment

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/apis/server"

	"github.com/labstack/echo/v4"
)

var _ = app.Router(NewRouter)

// NewArticle creates a new Article router object
func NewRouter(controllers IController) IRouter {
	return &Router{
		controller: controllers,
	}
}

// IArticle extends the IRouter interface
type IRouter interface {
	server.IRouter
}

// Article implements IArticle
type Router struct {
	controller IController
}

// Routers define route endpoints
func (r *Router) Routers(e *echo.Echo) {
	e.GET("/payment/status", r.controller.Status)
	e.POST("/payment", r.controller.Payment)
	e.POST("/payment/return", r.controller.PaymentReturn)
}
