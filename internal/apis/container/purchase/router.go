// Package purchase implements the router interface
// File purchase.go defines routes for APIs related to purchasing, adding to cart,
// accessing invoices, order, and creating invoices, order, etc.
package purchase

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/middleware"
	"swclabs/swix/internal/apis/server"

	"github.com/labstack/echo/v4"
)

var _ = app.Router(NewRouter)

// NewRouter returns a new Purchase router object
func NewRouter(controllers IController) IRouter {
	return &Router{controllers: controllers}
}

// IRouter extends the IRouter interface
type IRouter interface {
	server.IRouter
}

// Router is the router implementation for IPurchase
type Router struct {
	controllers IController
}

// Routers define route endpoint
func (p *Router) Routers(e *echo.Echo) {
	e.POST("/purchase/carts", p.controllers.AddToCarts, middleware.SessionProtected)
	e.GET("/purchase/carts", p.controllers.GetCarts)
	e.DELETE("/purchase/carts/:id", p.controllers.DeleteItem)

	e.GET("/purchase/orders", p.controllers.GetOrders)
	e.POST("/purchase/orders", p.controllers.CreateOrder, middleware.SessionProtected)

	e.GET("/address", p.controllers.GetDeliveryAddress)
	e.POST("/address", p.controllers.CreateDeliveryAddress)

	e.GET("/address/province", p.controllers.AddressProvince)
	e.GET("/address/district", p.controllers.AddressDistrict)
	e.GET("/address/ward", p.controllers.AddressWard)

	e.GET("/delivery", p.controllers.GetDelivery)
	e.POST("/delivery", p.controllers.CreateDelivery)
	e.POST("/delivery/order", p.controllers.CreateDeliveryOrder)
	e.GET("/delivery/order/:code", p.controllers.DeliveryOrderInfo)
}
