package router

import (
	"github.com/labstack/echo/v4"
	"swclabs/swipecore/internal/http/controller"
)

type Purchase struct {
	controllers controller.IPurchase
}

func (p Purchase) Routers(e *echo.Echo) {
	//TODO implement me
	panic("implement me")
}

func NewPurchase(controllers *controller.Purchase) *Purchase {
	return &Purchase{controllers: controllers}
}

var _ IRouter = (*Purchase)(nil)
