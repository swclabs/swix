package router

import (
	"github.com/labstack/echo/v4"
)

type IRouter interface {
	Routers(e *echo.Echo)
}
