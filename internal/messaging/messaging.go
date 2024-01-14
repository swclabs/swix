package messaging

import (
	"swclabs/swiftcart/internal/messaging/queue"
	"swclabs/swiftcart/internal/messaging/router"
	"swclabs/swiftcart/pkg/worker"
)

type Messaging struct {
	engine *worker.Engine
}

func NewMessaging() *Messaging {
	return &Messaging{
		engine: worker.NewServer(worker.Priority{
			queue.CriticalQueue: 6, // processed 60% of the time
			queue.DefaultQueue:  3, // processed 30% of the time
			queue.LowQueue:      1, // processed 10% of the time
		}),
	}
}

func (msg *Messaging) router(queue ...func(eng *worker.Engine)) {
	for _, q := range queue {
		q(msg.engine)
	}
}

func (msg *Messaging) Run(concurrency int) error {
	msg.router(router.Common)

	return msg.engine.Run(concurrency)
}
