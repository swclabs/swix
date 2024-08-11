package router

import (
	"swclabs/swix/internal/webapi/controller"

	"github.com/labstack/echo/v4"
)

// IPaydeliver interface for manager
type IPaydeliver interface {
	IRouter
}

// NewManager creates a new Manager router object
func NewPaydeliver(controllers controller.IPaydeliver) IPaydeliver {
	return &Paydeliver{
		controller: controllers,
	}
}

// Manager struct	implementation of IManager
type Paydeliver struct {
	controller controller.IPaydeliver
}

// Routers implements IPaydeliver.
func (p *Paydeliver) Routers(e *echo.Echo) {
	e.GET("/address", p.controller.GetDeliveryAddress)
	e.POST("/address", p.controller.CreateDeliveryAddress)

	e.GET("/delivery", p.controller.GetDelivery)
	e.POST("/delivery", p.controller.CreateDelivery)
}
