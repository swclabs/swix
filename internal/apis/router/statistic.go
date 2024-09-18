package router

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/controller"

	"github.com/labstack/echo/v4"
)

type IStatistic interface {
	IRouter
}

var _ IStatistic = (*Statistic)(nil)
var _ = app.Router(NewStatistic)

func NewStatistic(controller controller.IStatistic) IStatistic {
	return &Statistic{controller: controller}
}

type Statistic struct {
	controller controller.IStatistic
}

// Routers implements IStatistic.
func (s *Statistic) Routers(e *echo.Echo) {
	panic("unimplemented")
}
