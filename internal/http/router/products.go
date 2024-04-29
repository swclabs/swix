package router

import (
	"swclabs/swipecore/internal/http/controller"

	"github.com/labstack/echo/v4"
)

const TypeProducts = "products"

type Products struct {
	controller controller.IProducts
}

func newProducts() *Products {
	return &Products{
		controller: controller.NewProducts(),
	}
}

func (r *Products) Routers(e *echo.Echo) {
	e.GET("/categories", r.controller.GetCategories)
	e.GET("/products", r.controller.GetProductLimit)
	e.GET("/suppliers", r.controller.GetSupplier)
	e.POST("/suppliers", r.controller.InsertSupplier)
	e.POST("/categories", r.controller.InsertCategory)
	e.POST("/products", r.controller.UploadProduct)
	e.POST("/products/img", r.controller.UploadProductImage)
}
