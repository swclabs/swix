package middleware

import (
	"net/http"

	"github.com/swclabs/swipe-api/pkg/jwt"

	"github.com/gin-contrib/sessions"
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
	_, err := jwt.ParseToken(authHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":     "unauthorized",
			"success": false,
		})
		return
	}
	c.Next()
}

func SessionProtected(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("access_token") != nil {
		AccessToken := session.Get("access_token").(string)
		email, err := jwt.ParseToken(AccessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg":     "unauthorized",
				"success": false,
			})
			return
		}
		session.Set("email", email)
		if err := session.Save(); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":     "unauthorized",
			"success": false,
		})
	}
}
