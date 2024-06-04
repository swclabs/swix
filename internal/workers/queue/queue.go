package queue

import (
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/lib/worker"
)

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

func New() worker.Priority {
	return worker.Priority{
		CriticalQueue: 6, // processed 60% of the time
		DefaultQueue:  3, // processed 30% of the time
		LowQueue:      1, // processed 10% of the time
	}
}
