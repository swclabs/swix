package router

import (
	"swclabs/swipe-api/internal/http/controller"

	"github.com/labstack/echo/v4"
)

type Products struct {
	controller controller.IProducts
}

func NewProducts() *Products {
	return &Products{
		controller: controller.NewProducts(),
	}
}

func (r *Products) Common(e *echo.Echo) {
	e.GET("/newsletters", r.controller.GetNewsletter)
	e.GET("/categories", r.controller.GetCategories)
	e.GET("/products", r.controller.GetProductLimit)
}