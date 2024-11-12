package stats

import (
	"github.com/swclabs/swipex/internal/core/service/statistic"
)

var _ IController = (*Controller)(nil)

// var _ = sx.Controller(NewStatistic)

func NewController(service statistic.IStatistic) IController {
	return &Controller{service: service}
}

type IController interface {
}

type Controller struct {
	service statistic.IStatistic
}
