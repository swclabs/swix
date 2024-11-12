package stats

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/apis/server"

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
