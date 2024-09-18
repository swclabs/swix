package router

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/controller"

	"github.com/labstack/echo/v4"
)

var _ = app.Router(NewClassify)

// NewClassify returns a new Products router object
func NewClassify(controllers controller.IClassify) IClassify {
	return &Classify{
		controller: controllers,
	}
}

// IClassify router objects
type IClassify interface {
	IRouter
}

// Classify router objects
type Classify struct {
	controller controller.IClassify
}

// Routers implements IClassify.
func (c *Classify) Routers(e *echo.Echo) {
	// endpoint for suppliers
	e.GET("/suppliers", c.controller.GetSupplier)
	e.POST("/suppliers", c.controller.InsertSupplier)
	// TODO: implement edit supplier here
	// e.PUT("/suppliers")

	// endpoint for categories
	e.GET("/categories", c.controller.GetCategories)
	e.POST("/categories", c.controller.InsertCategory)
	e.DELETE("/categories", c.controller.DeleteCategory)
	e.PUT("/categories", c.controller.UpdateCategory)
}
