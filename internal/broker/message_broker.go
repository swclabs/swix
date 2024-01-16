package broker

import (
	"swclabs/swiftcart/internal/broker/queue"
	"swclabs/swiftcart/internal/broker/router"
	"swclabs/swiftcart/pkg/worker"
)

type Broker struct {
	engine *worker.Engine
}

func New() *Broker {
	return &Broker{
		engine: worker.NewServer(worker.Priority{
			queue.CriticalQueue: 6, // processed 60% of the time
			queue.DefaultQueue:  3, // processed 30% of the time
			queue.LowQueue:      1, // processed 10% of the time
		}),
	}
}

func (msg *Broker) router(queue ...func(eng *worker.Engine)) {
	for _, q := range queue {
		q(msg.engine)
	}
}

func (msg *Broker) Run(concurrency int) error {
	msg.router(router.Common)

	return msg.engine.Run(concurrency)
}
