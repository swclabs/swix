package stats

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/server"

	"github.com/labstack/echo/v4"
)

type IRouter interface {
	server.IRouter
}

var _ IRouter = (*Router)(nil)
var _ = app.Router(NewRouter)

func NewRouter(controller IController) IRouter {
	return &Router{controller: controller}
}

type Router struct {
	controller IController
}

// Routers implements IStatistic.
func (s *Router) Routers(e *echo.Echo) {
	panic("unimplemented")
}
