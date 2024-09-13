package statistic

var _ IStatistic = (*Statistic)(nil)

func NewStatistic() IStatistic {
	return &Statistic{}
}

type Statistic struct {
}
