package router

import "github.com/gin-gonic/gin"

type Customer struct{}

func (customer *Customer) Addresses(e *gin.Engine) {
	e.GET("/addresses")
	e.POST("/addresses")
	e.DELETE("/addresses/:id")
	e.PUT("/addresses/:id")
}

func (customer *Customer) Suppliers(e *gin.Engine) {
	e.GET("/suppliers")
	e.POST("/suppliers")
	e.PUT("/suppliers/:id")
	e.DELETE("/suppliers/:id")
}
