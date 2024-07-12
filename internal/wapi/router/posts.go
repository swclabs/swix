package router

import (
	"swclabs/swipecore/internal/wapi/controller"

	"github.com/labstack/echo/v4"
)

type IPosts interface {
	IRouter
}

type Posts struct {
	controller controller.IPosts
}

func NewPosts(controllers controller.IPosts) IPosts {
	return &Posts{
		controller: controllers,
	}
}

func (p *Posts) Routers(e *echo.Echo) {
	e.GET("/collections", p.controller.GetSlicesOfCollections)
	e.POST("/collections", p.controller.UploadCollections)
	e.PUT("/collections/img", p.controller.UpdateCollectionsImage)

	e.GET("/collections/headline", p.controller.GetSlicesOfHeadlineBanner)
	e.POST("/collections/headline", p.controller.UploadHeadlineBanner)
}
