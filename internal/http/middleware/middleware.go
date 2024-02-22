package middleware

import (
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/pkg/utils"
)

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

// func GinMiddleware(a *gin.Engine) {
// 	a.Use(CORSMiddleware())
// 	store := cookie.NewStore([]byte("secret"))
// 	a.Use(sessions.Sessions("mysession", store))
// }

func Sentry(e *echo.Echo) {
	if config.StageStatus != "dev" {
		e.Use(sentryecho.New(sentryecho.Options{
			Repanic:         true,
			WaitForDelivery: true,
		}))
	}
}

func CookieSetting(e *echo.Echo) {
	store := utils.NewSession()
	e.Use(session.Middleware(store))
}

func BaseSetting(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
