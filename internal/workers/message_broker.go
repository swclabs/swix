package workers

import (
	"swclabs/swipe-api/internal/workers/queue"
	"swclabs/swipe-api/internal/workers/router"
	"swclabs/swipe-api/pkg/tools/worker"
)

type Writer struct {
	engine *worker.Engine
}

func New() *Writer {
	return &Writer{
		engine: worker.NewServer(worker.Priority{
			queue.CriticalQueue: 6, // processed 60% of the time
			queue.DefaultQueue:  3, // processed 30% of the time
			queue.LowQueue:      1, // processed 10% of the time
		}),
	}
}

func (msg *Writer) router(queue ...func(eng *worker.Engine)) {
	for _, q := range queue {
		q(msg.engine)
	}
}

func (msg *Writer) Run(concurrency int) error {
	msg.router(
		router.Common,
		router.AccountManagement,
	)

	return msg.engine.Run(concurrency)
}
