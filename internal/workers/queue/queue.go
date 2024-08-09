// Package queue define queue name
package queue

import (
	"swclabs/swix/internal/config"
	"swclabs/swix/pkg/lib/worker"
)

var (
	// CriticalQueue define critical queue
	CriticalQueue = "critical"

	// DefaultQueue define default queue
	DefaultQueue = "default"

	// LowQueue define low queue
	LowQueue = "low"
)

var (
	// OrderQueue define order queue
	OrderQueue = "order"

	// CartQueue define cart queue
	CartQueue = "cart"
)

var _ = initQueue()

func initQueue() error {
	var queues = []*string{
		&CriticalQueue,
		&DefaultQueue,
		&LowQueue,

		&OrderQueue,
		&CartQueue,
	}

	if config.StageStatus != "prod" {
		for _, queue := range queues {
			*queue = *queue + "_dev"
		}
	}
	return nil
}

// New return worker priority
func New() worker.Priority {
	return worker.Priority{
		// CriticalQueue: 6, // processed 60% of the time
		// DefaultQueue:  3, // processed 30% of the time
		// LowQueue:      1, // processed 10% of the time
		CriticalQueue: 5, // processed 50% of the time
		CartQueue:     5, // processed 50% of the time
	}
}
