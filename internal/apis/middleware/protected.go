// Package middleware This file contains the middleware for protected routes.
package middleware

import (
	"net/http"
	"swclabs/swix/pkg/lib/crypto"

	"github.com/labstack/echo/v4"
)

// Protected middleware
func Protected(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// authHeader := c.Request().Header["Authorization"]
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"msg":     "unauthorized",
				"success": false,
			})
		}
		_, _, err := crypto.ParseToken(authHeader)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"msg":     "unauthorized",
				"success": false,
			})
		}
		return next(c)
	}
}
