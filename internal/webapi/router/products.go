// Package router implements the router interface
// File products.go defines APIs for functions related to products
// such as getting a list of products, adding a product, updating product information,
// deleting a product, adding product images.
// It also includes APIs related to inventory.
package router

import (
	"swclabs/swix/internal/webapi/controller"

	"github.com/labstack/echo/v4"
)

// IProducts router objects
type IProducts interface {
	IRouter
}

// Products router objects
type Products struct {
	controller controller.IProducts
}

// NewProducts returns a new Products router object
func NewProducts(controllers controller.IProducts) IProducts {
	return &Products{
		controller: controllers,
	}
}

// Routers define route endpoints
func (r *Products) Routers(e *echo.Echo) {
	// endpoint for search
	e.GET("/search", r.controller.Search)

	// endpoint for products
	e.GET("/products", r.controller.GetProductLimit)
	e.POST("/products", r.controller.CreateProduct)
	e.PUT("/products", r.controller.UpdateProductInfo)
	e.DELETE("/products", r.controller.DeleteProduct)
	e.GET("/products/details", r.controller.GetProductDetails)
	e.GET("/products/accessory", r.controller.AccessoryDetail)
	e.POST("/products/img", r.controller.UploadProductImage)
	e.GET("/products/:type", r.controller.GetProductView)

	// endpoint for inventories
	e.GET("/inventories", r.controller.GetStock)
	e.PUT("/inventories", r.controller.UpdateInv)
	e.DELETE("/inventories", r.controller.DeleteInv)
	e.PUT("/inventories/image", r.controller.UploadInvImage)
	e.GET("/inventories/details", r.controller.GetInvDetails)
	e.POST("/inventories/:type", r.controller.AddInv)
	e.POST("/inventories/specs/:type", r.controller.InsertInvSpecs)
}
