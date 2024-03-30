package middleware

import (
	"net/http"
	"os"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"swclabs/swipe-api/internal/config"
	"swclabs/swipe-api/pkg/utils"
)

func CORS() echo.MiddlewareFunc {
	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	return middleware.CORSWithConfig(DefaultCORSConfig)
}

func Sentry(e *echo.Echo) {
	if config.StageStatus != "dev" {
		e.Use(sentryecho.New(sentryecho.Options{
			Repanic:         true,
			WaitForDelivery: true,
		}))
	}
}

func Logger(file *os.File, e *echo.Echo) {

	conf := middleware.DefaultLoggerConfig
	conf.Output = file

	e.Use(middleware.LoggerWithConfig(conf))
}

func CookieSetting(e *echo.Echo) {
	store := utils.NewSession()
	e.Use(session.Middleware(store))
}

func BaseSetting(e *echo.Echo) {
	// accept any domain
	e.Use(CORS())
	// use logger to write logs to api.log file
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
