package router

import (
	_ "swclabs/swipe-api/docs"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

const TypeDocs = "docs"

type Docs struct{}

func newDocs() *Docs {
	return &Docs{}
}

// Docs
// API documentation
// Router documentation
// Base on: http://${HOST}:${PORT}/docs/index.html#/
func (d *Docs) Routers(e *echo.Echo) {
	e.GET("/docs/*any", echoSwagger.WrapHandler)
}
