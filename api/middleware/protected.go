package middleware

import (
	"example/swiftcart/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Protected(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":     "unauthorized",
			"success": false,
		})
		return
	}
	_, err := utils.ParseToken(authHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":     "unauthorized",
			"success": false,
		})
		return
	}
	c.Next()
}
