package router

import (
	"swclabs/swipecore/internal/http/controller"

	"github.com/labstack/echo/v4"
)

const TypePosts = "posts"

type Posts struct {
	controller controller.IPosts
}

func newPosts() *Posts {
	return &Posts{
		controller: controller.NewPosts(),
	}
}

func (p *Posts) Routers(e *echo.Echo) {
	e.GET("/newsletters", p.controller.GetNewsletter)
	e.POST("/newsletters", p.controller.UploadNewsletter)
}
