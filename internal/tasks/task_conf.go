package tasks

import (
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/x/worker"
)

var (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
)

const (
	WorkerHealthCheck string = "Worker.HealthCheck"
)

func init() {
	if config.StageStatus != "prod" {
		CriticalQueue = "critical_dev"
		DefaultQueue = "default_dev"
		LowQueue = "low_dev"
	}
}

func Queue() worker.Queue {
	return worker.Queue{
		CriticalQueue: 6, // processed 60% of the time
		DefaultQueue:  3, // processed 30% of the time
		LowQueue:      1, // processed 10% of the time
	}
}

func Path() worker.Path {
	return worker.Path{
		WorkerHealthCheck: HandleHealthCheck,
	}
}
