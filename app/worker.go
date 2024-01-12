package app

import (
	"swclabs/swiftcart/delivery/messaging"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/worker"
)

func init() {
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

type IWorker interface {
	Run() error
}

type Worker struct {
	engine *worker.Engine
}

func NewWorker(concurrency int) IWorker {
	return &Worker{
		engine: worker.NewServer(concurrency, messaging.Queue()),
	}
}

func (w *Worker) Run() error {
	w.engine.HandleFunctions(messaging.Path())
	return w.engine.Run()
}
