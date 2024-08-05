// Package middleware This file contains the middleware for protected routes.
package middleware

import (
	"net/http"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/pkg/lib/crypto"

	"swclabs/swix/pkg/utils"

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
		_, err := crypto.ParseToken(authHeader)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"msg":     "unauthorized",
				"success": false,
			})
		}
		return next(c)
	}
}

// SessionProtected middleware
func SessionProtected(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// session := sessions.Default(c)
		AccessToken := utils.Session(c, utils.BaseSessions, "access_token")
		if AccessToken != nil {
			email, err := crypto.ParseToken(AccessToken.(string))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"msg":     "unauthorized",
					"success": false,
				})
			}
			if err := utils.SaveSession(c, utils.BaseSessions, "email", email); err != nil {
				return c.JSON(http.StatusInternalServerError, dtos.Error{
					Msg: err.Error(),
				})
			}
			role, _ := crypto.ParseTokenRole(AccessToken.(string))
			if role != "" {
				if err := utils.SaveSession(c, utils.BaseSessions, "role", role); err != nil {
					return c.JSON(http.StatusInternalServerError, dtos.Error{
						Msg: err.Error(),
					})
				}
			}
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"msg":     "unauthorized",
			"success": false,
		})
	}
}

// RequireAdmin middleware for admin routes
func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// session := sessions.Default(c)
		AccessToken := utils.Session(c, utils.BaseSessions, "access_token")
		if AccessToken != nil {
			email, err := crypto.ParseToken(AccessToken.(string))
			if err != nil {
				return c.Redirect(http.StatusSeeOther, "/")
			}
			if err := utils.SaveSession(c, utils.BaseSessions, "email", email); err != nil {
				return c.Redirect(http.StatusSeeOther, "/")
			}
			role, _ := crypto.ParseTokenRole(AccessToken.(string))
			if role != "" {
				if err := utils.SaveSession(c, utils.BaseSessions, "role", role); err != nil {
					return c.Redirect(http.StatusSeeOther, "/")
				}
			} else {
				return c.Redirect(http.StatusSeeOther, "/")
			}
			return next(c)
		}
		return c.Redirect(http.StatusSeeOther, "/")
	}
}
