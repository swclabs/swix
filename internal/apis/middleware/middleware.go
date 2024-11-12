// Package middleware define middleware
package middleware

import (
	"net/http"
	"os"

	"github.com/swclabs/swipex/internal/config"

	sessionUtils "github.com/swclabs/swipex/pkg/lib/session"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORS middleware
func CORS() echo.MiddlewareFunc {
	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	return middleware.CORSWithConfig(DefaultCORSConfig)
}

// Sentry middleware
func Sentry(e *echo.Echo) {
	if config.StageStatus != "dev" {
		e.Use(sentryecho.New(sentryecho.Options{
			Repanic:         true,
			WaitForDelivery: true,
		}))
	}
}

// Logger middleware
func Logger(file *os.File, e *echo.Echo) {

	conf := middleware.DefaultLoggerConfig
	conf.Output = file

	e.Use(middleware.LoggerWithConfig(conf))
}

// CookieSetting middleware
func CookieSetting(e *echo.Echo) {
	store := sessionUtils.New()
	e.Use(session.Middleware(store))
}

// BaseSetting middleware
func BaseSetting(e *echo.Echo) {
	// accept any domain
	e.Use(CORS())
	// use logger to write logs to api.log file
	e.Use(middleware.Logger())
	if config.StageStatus != "dev" {
		e.Use(middleware.Recover())
	}
}
