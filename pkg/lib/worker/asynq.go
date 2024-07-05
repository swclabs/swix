package worker

import (
	"context"
	"fmt"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/lib/logger"

	"github.com/hibiken/asynq"
)

type (
	// Priority is a queue priority define in asynq
	Priority map[string]int

	// Queue is a map of queue name and handler function
	Queue map[string]func(context.Context, *asynq.Task) error

	// HandleFunc is a sort of handler function type
	HandleFunc func(ctx context.Context, task *asynq.Task) error
)

// Engine includes all components of the asynq server
type Engine struct {
	server   *asynq.Server
	mux      *asynq.ServeMux
	priority Priority
	queue    Queue
	broker   asynq.RedisClientOpt
}

func New(priorityQueue Priority) *Engine {
	return &Engine{
		server:   nil,
		mux:      asynq.NewServeMux(),
		queue:    Queue{},
		priority: priorityQueue,
		broker: asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort), // Redis server address
			Password: config.RedisPassword,                                     // Redis password
		},
	}
}

// handleFunctions run all functions in the given path
func (w *Engine) handleFunctions() {
	for k, v := range w.queue {
		w.mux.HandleFunc(k, v)
	}
}

func (w *Engine) RegisterQueue(hfn func() (taskName string, fn HandleFunc)) {
	taskName, fn := hfn()
	w.queue[taskName] = fn
}

func (w *Engine) Run(concurrency int) error {
	// Create a new Asynq server
	w.server = asynq.NewServer(w.broker, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: concurrency,
		// Specify multiple queues with different priority.
		Queues: w.priority,
	})
	w.handleFunctions()

	logger.Banner("Launching a asynchronous worker with the following settings:")
	logger.Broker("redis", w.broker.Addr)
	for q, p := range w.priority {
		logger.Queue(q, p)
	}
	logger.Banner("Handle Function: ")
	for types, handler := range w.queue {
		logger.HandleFunc(types, getName(handler))
	}
	fmt.Println()
	return w.server.Run(w.mux)
}
