package controller

import (
	"swclabs/swix/internal/core/service/statistic"
)

type IStatistic interface {
}

var _ IStatistic = (*Statistic)(nil)

func NewStatistic(service statistic.IStatistic) IStatistic {
	return &Statistic{service: service}
}

type Statistic struct {
	service statistic.IStatistic
}
