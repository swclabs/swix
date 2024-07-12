// Package router implements the router interface
package router

import (
	"swclabs/swipecore/internal/wapi/controller"

	"github.com/labstack/echo/v4"
)

// IPurchase extend the IRouter interface
type IPurchase interface {
	IRouter
}

// Purchase router implementation IPurchase
type Purchase struct {
	controllers controller.IPurchase
}

var _ IRouter = (*Purchase)(nil)

// NewPurchase returns a new Purchase router object
func NewPurchase(controllers controller.IPurchase) IPurchase {
	return &Purchase{controllers: controllers}
}

// Routers define route endpoint
func (p *Purchase) Routers(e *echo.Echo) {
	e.POST("/purchase/carts", p.controllers.AddToCarts)
	e.GET("/purchase/carts", p.controllers.GetCarts)
}
