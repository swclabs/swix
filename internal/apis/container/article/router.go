package article

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/server"

	"github.com/labstack/echo/v4"
)

var _ = app.Router(NewRouter)

// NewArticle creates a new Article router object
func NewRouter(controllers IController) IRouter {
	return &Router{
		controller: controllers,
	}
}

// IArticle extends the IRouter interface
type IRouter interface {
	server.IRouter
}

// Article implements IArticle
type Router struct {
	controller IController
}

// Routers define route endpoints
func (r *Router) Routers(e *echo.Echo) {
	e.GET("/news", r.controller.GetNews)
	e.POST("/news", r.controller.UploadNews)
	e.PUT("/news/img", r.controller.UpdateNewsImage)

	e.GET("/comment", r.controller.GetComment)
	e.POST("/comment", r.controller.UploadComment)
}
