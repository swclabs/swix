package queue

import "swclabs/swipecore/internal/config"

var (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
)

var (
	OrderQueue = "order"
	Purchase   = "purchase"
)

func init() {
	var queues = []*string{
		&CriticalQueue,
		&DefaultQueue,
		&LowQueue,

		&OrderQueue,
		&Purchase,
	}

	if config.StageStatus != "prod" {
		for _, queue := range queues {
			*queue = *queue + "_dev"
		}
	}
}
