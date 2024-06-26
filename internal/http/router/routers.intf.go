package router

import (
	"swclabs/swipecore/internal/http/controller"

	"github.com/labstack/echo/v4"
)

type IRouter interface {
	Routers(e *echo.Echo)
}

func New(common controller.ICommon) []IRouter {
	return []IRouter{
		NewDocs(),
		NewCommon(common),
	}
}
