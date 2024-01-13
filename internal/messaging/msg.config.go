package messaging

import (
	"swclabs/swiftcart/internal/messaging/handler"
	"swclabs/swiftcart/internal/messaging/queue"
	"swclabs/swiftcart/pkg/worker"
)

// Controller : example
//
//	return worker.Path{
//		WorkerHealthCheck: HandleHealthCheck,
//	}
func Controller() worker.Path {
	return worker.GetPath(
		handler.CommonPath(), // common path
	)
}

func Queue() worker.Queue {
	return worker.Queue{
		queue.CriticalQueue: 6, // processed 60% of the time
		queue.DefaultQueue:  3, // processed 30% of the time
		queue.LowQueue:      1, // processed 10% of the time
	}
}
