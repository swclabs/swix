package router

import "github.com/labstack/echo/v4"

type Customer struct{}

func (customer *Customer) Addresses(e *echo.Echo) {
	// e.GET("/addresses")
	// e.POST("/addresses")
	// e.DELETE("/addresses/:id")
	// e.PUT("/addresses/:id")
}

func (customer *Customer) Suppliers(e *echo.Echo) {
	// e.GET("/suppliers")
	// e.POST("/suppliers")
	// e.PUT("/suppliers/:id")
	// e.DELETE("/suppliers/:id")
}
