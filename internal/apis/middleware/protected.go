// Package middleware This file contains the middleware for protected routes.
package middleware

import (
	"fmt"
	"net/http"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/pkg/lib/crypto"
	"swclabs/swix/pkg/lib/session"

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

// SessionProtected middleware
func SessionProtected(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// session := sessions.Default(c)
		AccessToken := session.Get(c, session.Base, "access_token")
		if AccessToken != "" {
			accountID, email, err := crypto.ParseToken(AccessToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"msg":     "unauthorized",
					"success": false,
				})
			}
			if err := session.Save(c, session.Base, "email", email); err != nil {
				return c.JSON(http.StatusInternalServerError, dtos.Error{
					Msg: err.Error(),
				})
			}
			if err := session.Save(c, session.Base, "account_id", fmt.Sprintf("%d", accountID)); err != nil {
				return c.JSON(http.StatusInternalServerError, dtos.Error{
					Msg: err.Error(),
				})
			}
			role, _ := crypto.ParseTokenRole(AccessToken)
			if role != "" {
				if err := session.Save(c, session.Base, "role", role); err != nil {
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
		AccessToken := session.Get(c, session.Base, "access_token")
		if AccessToken != "" {
			_, email, err := crypto.ParseToken(AccessToken)
			if err != nil {
				return c.Redirect(http.StatusSeeOther, "/")
			}
			if err := session.Save(c, session.Base, "email", email); err != nil {
				return c.Redirect(http.StatusSeeOther, "/")
			}
			role, _ := crypto.ParseTokenRole(AccessToken)
			if role != "" {
				if err := session.Save(c, session.Base, "role", role); err != nil {
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
