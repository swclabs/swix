package delivery

import (
	"swclabs/swiftcart/internal/broker"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/worker"
)

func init() {
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

type IWorker interface {
	Run(concurrency int) error
}

type Worker struct {
	engine *broker.Broker
}

func NewWorker() IWorker {
	return &Worker{
		engine: broker.New(),
	}
}

func (w *Worker) Run(concurrency int) error {
	return w.engine.Run(concurrency)
}
