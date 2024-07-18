// Package router implements the router interface
// The file posts.go contains APIs related to getting a list of posts,
// adding a post, updating post information, deleting a post, and adding post images.
package router

import (
	"swclabs/swipecore/internal/webapi/controller"

	"github.com/labstack/echo/v4"
)

// IPosts extends the IRouter interface
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

// Routers define route endpoints
func (p *Posts) Routers(e *echo.Echo) {
	e.GET("/collections", p.controller.GetSlicesOfCollections)
	e.POST("/collections", p.controller.UploadCollections)
	e.PUT("/collections/img", p.controller.UpdateCollectionsImage)

	e.GET("/collections/headline", p.controller.GetSlicesOfHeadlineBanner)
	e.POST("/collections/headline", p.controller.UploadHeadlineBanner)
}
