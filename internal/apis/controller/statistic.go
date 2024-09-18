package controller

import (
	"swclabs/swix/internal/core/service/statistic"
)

var _ IStatistic = (*Statistic)(nil)

// var _ = sx.Controller(NewStatistic)

func NewStatistic(service statistic.IStatistic) IStatistic {
	return &Statistic{service: service}
}

type IStatistic interface {
}

type Statistic struct {
	service statistic.IStatistic
}
