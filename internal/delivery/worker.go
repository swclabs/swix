package delivery

import (
	"github.com/swclabs/swipe-api/internal/broker"
	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/pkg/worker"
)

type IWorker interface {
	Run(concurrency int) error
}

type _Worker struct {
	engine *broker.Broker
}

func NewWorker() IWorker {
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
	return &_Worker{
		engine: broker.New(),
	}
}

func (w *_Worker) Run(concurrency int) error {
	return w.engine.Run(concurrency)
}
