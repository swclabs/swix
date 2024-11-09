// Package purchase implements the router interface
// File purchase.go defines routes for APIs related to purchasing, adding to cart,
// accessing invoices, order, and creating invoices, order, etc.
package purchase

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/middleware"
	"swclabs/swipex/internal/apis/server"

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
	e.GET("/purchase/carts", p.controllers.GetCarts, middleware.Protected)
	e.POST("/purchase/carts", p.controllers.AddToCarts, middleware.Protected)
	e.DELETE("/purchase/carts/:id", p.controllers.DeleteCartItem, middleware.Protected)

	e.GET("/purchase/orders", p.controllers.GetOrders, middleware.Protected)
	e.GET("/purchase/orders/:code", p.controllers.GetOrdersByCode)
	e.POST("/purchase/orders", p.controllers.CreateOrder, middleware.Protected)

	e.GET("/purchase/admin/orders", p.controllers.GetOrdersByAdmin)
	e.POST("/purchase/admin/orders", p.controllers.CreateOrderForm)

	e.GET("/purchase/coupons", p.controllers.GetCoupon)
	e.GET("/purchase/coupons/:code", p.controllers.UseCoupon, middleware.Protected)
	e.POST("/purchase/coupons", p.controllers.CreateCoupon)

	e.GET("/address", p.controllers.GetDeliveryAddress, middleware.Protected)
	e.POST("/address", p.controllers.CreateDeliveryAddress)

	e.GET("/address/province", p.controllers.AddressProvince)
	e.GET("/address/district", p.controllers.AddressDistrict)
	e.GET("/address/ward", p.controllers.AddressWard)

	e.GET("/delivery", p.controllers.GetDelivery, middleware.Protected)
	e.GET("/delivery/order/:code", p.controllers.DeliveryOrderInfo)
	e.POST("/delivery", p.controllers.CreateDelivery)
	e.POST("/delivery/order", p.controllers.CreateDeliveryOrder)
}
