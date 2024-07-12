// Package router implements the router interface
package router

import (
	_ "swclabs/swipecore/docs" // init swagger docs

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// IDocs extend the IRouter interface
type IDocs interface {
	IRouter
}

// Docs implement IDocs interface
type Docs struct{}

// NewDocs creates a new Docs router object
func NewDocs() IDocs {
	return &Docs{}
}

// Routers .
// API documentation
// Register documentation
// Base on: http://${HOST}:${PORT}/docs/index.html#/
func (d *Docs) Routers(e *echo.Echo) {
	e.GET("/docs/*any", echoSwagger.WrapHandler)
}
