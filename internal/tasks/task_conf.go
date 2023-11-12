package tasks

import (
	"example/komposervice/internal/config"
	"example/komposervice/pkg/lib/worker"
)

var (
	WorkerQueue   = "default_queue"
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
)

func init() {
	if config.StageStatus != "prod" {
		WorkerQueue = "default_queue_dev"
		CriticalQueue = "critical_dev"
		DefaultQueue = "default_dev"
		LowQueue = "low_dev"
	}
}

const (
	WorkerHealthCheck string = "Worker.HealthCheck"
	WorkerSaveUser    string = "Worker.SaveUser"
)

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
