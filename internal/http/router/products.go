package router

import (
	"swclabs/swipecore/internal/http/controller"

	"github.com/labstack/echo/v4"
)

type IProducts interface {
	IRouter
}

type Products struct {
	controller controller.IProducts
}

func NewProducts(controllers controller.IProducts) IProducts {
	return &Products{
		controller: controllers,
	}
}

func (r *Products) Routers(e *echo.Echo) {
	// endpoint for category
	e.GET("/categories", r.controller.GetCategories)
	e.POST("/categories", r.controller.InsertCategory)
	// endpoint for products
	e.GET("/products", r.controller.GetProductLimit)
	e.POST("/products", r.controller.CreateProduct)
	e.DELETE("/products", r.controller.DeleteProduct)
	e.PUT("/products", r.controller.UpdateProductInfo)
	e.POST("/products/img", r.controller.UploadProductImage)
	// endpoint for suppliers
	e.GET("/suppliers", r.controller.GetSupplier)
	e.POST("/suppliers", r.controller.InsertSupplier)
	// endpoint for inventories
	e.GET("/inventories", r.controller.GetProductAvailability)
	e.POST("/inventories", r.controller.AddToInventory)
}
