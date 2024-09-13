package router

import (
	"swclabs/swix/internal/webapi/controller"

	"github.com/labstack/echo/v4"
)

type IStatistic interface {
	IRouter
}

var _ IStatistic = (*Statistic)(nil)

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
