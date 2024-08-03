// Package router implements the router interface
// The file article.go contains APIs related to getting a list of article,
// adding a post, updating post information, deleting a post, and adding post images.
package router

import (
	"swclabs/swipecore/internal/webapi/controller"

	"github.com/labstack/echo/v4"
)

// IArticle extends the IRouter interface
type IArticle interface {
	IRouter
}

// Article implements IArticle
type Article struct {
	controller controller.IArticle
}

// NewArticle creates a new Article router object
func NewArticle(controllers controller.IArticle) IArticle {
	return &Article{
		controller: controllers,
	}
}

// Routers define route endpoints
func (p *Article) Routers(e *echo.Echo) {
	e.GET("/collections", p.controller.GetArticleData)
	e.POST("/collections", p.controller.UploadArticle)
	e.PUT("/collections/img", p.controller.UpdateCollectionsImage)

	e.GET("/collections/headline", p.controller.GetSlicesOfHeadlineBanner)
	e.POST("/collections/headline", p.controller.UploadHeadlineBanner)
}
