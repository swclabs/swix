package router

import (
	_ "github.com/swclabs/swipe-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// API documentation
// Router documentation
// Base on: http://${HOST}:${PORT}/docs/index.html#/
func Docs(e *gin.Engine) {
	e.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
