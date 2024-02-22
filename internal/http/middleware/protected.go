package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/pkg/tools"
	"github.com/swclabs/swipe-api/pkg/utils"
)

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
		_, err := tools.ParseToken(authHeader)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"msg":     "unauthorized",
				"success": false,
			})
		}
		return next(c)
	}
}

func SessionProtected(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// session := sessions.Default(c)
		AccessToken := utils.Session(c, utils.BaseSessions, "access_token")
		log.Println(AccessToken)
		if AccessToken != nil {
			email, err := tools.ParseToken(AccessToken.(string))
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"msg":     "unauthorized",
					"success": false,
				})
			}
			// session.Set("email", email)
			// if err := session.Save(); err != nil {
			// 	return c.String(http.StatusInternalServerError, err.Error())
			// }
			if err := utils.SaveSession(c, utils.BaseSessions, "email", email); err != nil {
				return c.JSON(http.StatusInternalServerError, domain.Error{
					Msg: err.Error(),
				})
			}
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"msg":     "unauthorized",
				"success": false,
			})
		}
		return next(c)
	}
}
