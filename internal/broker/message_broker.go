package broker

import (
	"github.com/swclabs/swipe-server/internal/broker/queue"
	"github.com/swclabs/swipe-server/internal/broker/router"
	"github.com/swclabs/swipe-server/pkg/tools/worker"
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
	msg.router(
		router.Common,
		router.AccountManagement,
	)

	return msg.engine.Run(concurrency)
}
