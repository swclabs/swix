package router

import "github.com/gin-gonic/gin"

type ProductManagement struct {
}

func (product *ProductManagement) Product(e *gin.Engine) {
	router := e.Group("/v1")
	router.POST("/products")
	router.GET("/products")
	router.PUT("/products/:id")
	router.DELETE("/products/:id")
}

func (product *ProductManagement) Category(e *gin.Engine) {
	router := e.Group("/v1")
	router.POST("/categories")
	router.GET("/categories")
	router.PUT("/categories/:id")
	router.DELETE("/categories/:id")
}
