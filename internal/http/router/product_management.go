package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swclabs/swipe-api/internal/http/controller"
)

type ProductManagement struct {
	controller controller.IProductManagement
}

func NewProductManagement() *ProductManagement {
	return &ProductManagement{
		controller: controller.NewProductManagement(),
	}
}

func (product *ProductManagement) Product(e *gin.Engine) {
	e.POST("/products")
	e.GET("/products")
	e.PUT("/products/:id")
	e.DELETE("/products/:id")
}

func (product *ProductManagement) Category(e *gin.Engine) {
	e.POST("/categories", product.controller.InsertCategory)
	// router.GET("/categories")
	// router.PUT("/categories/:id")
	// router.DELETE("/categories/:id")
}
