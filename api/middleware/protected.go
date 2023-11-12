package middleware

import (
	"example/komposervice/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Protected(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"msg":     "unauthorized",
			"success": false,
		})
		return
	}
	_, err := utils.ParseToken(authHeader)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"msg":     "unauthorized",
			"success": false,
		})
		return
	}
	c.Next()
}
