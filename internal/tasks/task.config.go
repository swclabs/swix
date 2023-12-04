package tasks

import (
	"maps"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/x/worker"
)

var (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"
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

func WorkerGetPath(path ...worker.Path) worker.Path {
	workerPath := worker.Path{}
	for _, p := range path {
		if p != nil {
			maps.Copy(workerPath, p)
		}
	}
	return workerPath
}

// Path : example
//
//	return worker.Path{
//		WorkerHealthCheck: HandleHealthCheck,
//	}
func Path() worker.Path {
	return WorkerGetPath(
		// common path
		worker.Path{
			WorkerHealthCheck: HandleHealthCheck,
		},
		// account management path
		AccountManagementPath(),
	)
}
