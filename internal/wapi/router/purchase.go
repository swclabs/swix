package router

import (
	"swclabs/swipecore/internal/wapi/controller"

	"github.com/labstack/echo/v4"
)

type IPurchase interface {
	IRouter
}

type Purchase struct {
	controllers controller.IPurchase
}

var _ IRouter = (*Purchase)(nil)

func NewPurchase(controllers controller.IPurchase) IPurchase {
	return &Purchase{controllers: controllers}
}

func (p *Purchase) Routers(e *echo.Echo) {
	e.POST("/purchase/carts", p.controllers.AddToCarts)
	e.GET("/purchase/carts", p.controllers.GetCarts)
}
