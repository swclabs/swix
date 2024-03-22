package delivery

import (
	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/internal/workers"
	"github.com/swclabs/swipe-api/pkg/tools/worker"
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
