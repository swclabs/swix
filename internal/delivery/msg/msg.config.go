package msg

import (
	"swclabs/swiftcart/internal/delivery/msg/handler"
	"swclabs/swiftcart/internal/delivery/msg/queue"
	"swclabs/swiftcart/pkg/worker"
)

// Path : example
//
//	return worker.Path{
//		WorkerHealthCheck: HandleHealthCheck,
//	}
func Path() worker.Path {
	return worker.GetPath(
		// common path
		handler.CommonPath(),
	)
}

func Queue() worker.Queue {
	return worker.Queue{
		queue.CriticalQueue: 6, // processed 60% of the time
		queue.DefaultQueue:  3, // processed 30% of the time
		queue.LowQueue:      1, // processed 10% of the time
	}
}
