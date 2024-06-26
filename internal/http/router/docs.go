package router

import (
	_ "swclabs/swipecore/docs"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

const TypeDocs = "docs"

type Docs struct{}

func NewDocs() IRouter {
	return &Docs{}
}

// Routers
// API documentation
// Register documentation
// Base on: http://${HOST}:${PORT}/docs/index.html#/
func (d *Docs) Routers(e *echo.Echo) {
	e.GET("/docs/*any", echoSwagger.WrapHandler)
}
