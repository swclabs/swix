package boot

import (
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/workers"
	"swclabs/swipecore/pkg/tools/worker"
)

type IWorker interface {
	Run(concurrency int) error
}

type _Worker struct {
	engine *workers.Writer
}

func NewWorker() IWorker {
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
	return &_Worker{
		engine: workers.New(),
	}
}

func (w *_Worker) Run(concurrency int) error {
	return w.engine.Run(concurrency)
}
