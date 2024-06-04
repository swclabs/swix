package router

import (
	"swclabs/swipecore/internal/http/controller"

	"github.com/labstack/echo/v4"
)

const TypePosts = "posts"

type Posts struct {
	controller controller.IPosts
}

func NewPosts(controllers *controller.Posts) *Posts {
	return &Posts{
		controller: controllers,
	}
}

func (p *Posts) Routers(e *echo.Echo) {
	e.POST("/collections", p.controller.UploadCollections)
	e.PUT("/collections/img", p.controller.UpdateCollectionsImage)
	e.GET("/collections", p.controller.GetSlicesOfCollections)
}
