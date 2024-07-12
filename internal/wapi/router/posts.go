// Package router implements the router interface
package router

import (
	"swclabs/swipecore/internal/wapi/controller"

	"github.com/labstack/echo/v4"
)

// IPosts extends IRouter interface
type IPosts interface {
	IRouter
}

// Posts implements IPosts
type Posts struct {
	controller controller.IPosts
}

// NewPosts creates a new Posts router object
func NewPosts(controllers controller.IPosts) IPosts {
	return &Posts{
		controller: controllers,
	}
}

// Routers define route endpoint
func (p *Posts) Routers(e *echo.Echo) {
	e.GET("/collections", p.controller.GetSlicesOfCollections)
	e.POST("/collections", p.controller.UploadCollections)
	e.PUT("/collections/img", p.controller.UpdateCollectionsImage)

	e.GET("/collections/headline", p.controller.GetSlicesOfHeadlineBanner)
	e.POST("/collections/headline", p.controller.UploadHeadlineBanner)
}
