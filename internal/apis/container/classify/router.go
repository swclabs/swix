package classify

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/server"

	"github.com/labstack/echo/v4"
)

var _ = app.Router(NewRouter)

// NewClassify returns a new Products router object
func NewRouter(controllers IController) IRouter {
	return &Router{
		controller: controllers,
	}
}

// IClassify router objects
type IRouter interface {
	server.IRouter
}

// Classify router objects
type Router struct {
	controller IController
}

// Routers implements IClassify.
func (r *Router) Routers(e *echo.Echo) {
	// endpoint for suppliers
	e.GET("/suppliers", r.controller.GetSupplier)
	e.POST("/suppliers", r.controller.InsertSupplier)
	// TODO: implement edit supplier here
	// e.PUT("/suppliers")

	// endpoint for categories
	e.GET("/categories", r.controller.GetCategories)
	e.POST("/categories", r.controller.InsertCategory)
	e.DELETE("/categories", r.controller.DeleteCategory)
	e.PUT("/categories", r.controller.UpdateCategory)
}
