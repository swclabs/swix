package delivery

import (
	"swclabs/swiftcart/internal/broker"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/worker"
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
