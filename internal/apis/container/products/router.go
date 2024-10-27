// Package products implements the router interface
// File products.go defines APIs for functions related to products
// such as getting a list of products, adding a product, updating product information,
// deleting a product, adding product images.
// It also includes APIs related to inventory.
package products

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/middleware"
	"swclabs/swipex/internal/apis/server"

	"github.com/labstack/echo/v4"
)

var _ = app.Router(NewRouter)

// NewRouter returns a new Products router object
func NewRouter(controllers IController) IRouter {
	return &Router{
		controller: controllers,
	}
}

// IRouter router objects
type IRouter interface {
	server.IRouter
}

// Router router objects
type Router struct {
	controller IController
}

// Routers define route endpoints
func (r *Router) Routers(e *echo.Echo) {
	// endpoint for search
	e.GET("/search", r.controller.Search)
	e.GET("/search/details", r.controller.SearchDetails)

	// endpoint for products
	e.GET("/products", r.controller.GetProductLimit)
	e.POST("/products", r.controller.CreateProduct)
	e.PUT("/products", r.controller.UpdateProductInfo)
	e.DELETE("/products", r.controller.DeleteProduct)
	e.GET("/products/details", r.controller.GetProductDetails)
	e.PUT("/products/img", r.controller.UploadProductImage)
	e.PUT("/products/img/shop", r.controller.UploadProductShopImage)
	e.GET("/products/:type", r.controller.GetProductByType)
	e.PUT("/rating/:id", r.controller.Rating, middleware.Protected)

	// endpoint for inventories
	e.GET("/inventories", r.controller.GetItems)
	e.PUT("/inventories", r.controller.UpdateInv)
	e.DELETE("/inventories", r.controller.DeleteInv)
	e.PUT("/inventories/image", r.controller.UploadInvImage)
	e.PUT("/inventories/image/color", r.controller.UploadInvColorImage)
	e.GET("/inventories/details", r.controller.GetInvDetails)
	e.POST("/inventories", r.controller.InsertInv)
}
