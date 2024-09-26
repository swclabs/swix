// Package worker define worker writer engine
package worker

import (
	"context"
	"fmt"
	"strconv"
	"swclabs/swix/internal/config"
	"swclabs/swix/pkg/lib/logger"

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

// IEngine is an interface for Engine
type IEngine interface {
	Register(taskName string, fn HandleFunc)
	HandlerFunc(hfn func() (taskName string, fn HandleFunc))
	Run(concurrency int) error
}

// Engine includes all components of the asynq server
type Engine struct {
	server   *asynq.Server
	mux      *asynq.ServeMux
	priority Priority
	queue    Queue
	broker   asynq.RedisClientOpt
}

// New creates a new instance of the Engine
func New(priorityQueue Priority) IEngine {
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

// Register implements IEngine.
func (w *Engine) Register(taskName string, fn HandleFunc) {
	w.queue[taskName] = fn
}

// handleFunctions run all functions in the given path
func (w *Engine) handleFunctions() {
	for tn, fn := range w.queue {
		w.mux.HandleFunc(tn, fn)
	}
}

// HandlerFunc register the queue
func (w *Engine) HandlerFunc(hfn func() (taskName string, fn HandleFunc)) {
	taskName, fn := hfn()
	w.queue[taskName] = fn
}

// Run start the asynq server
func (w *Engine) Run(concurrency int) error {
	// Create a new Asynq server
	w.server = asynq.NewServer(w.broker, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: concurrency,
		// Specify multiple queues with different priority.
		Queues: w.priority,
	})
	w.handleFunctions()

	logger.Info("Launching a asynchronous worker with the following settings:")
	logger.Info("Redis: " + w.broker.Addr)
	for q, p := range w.priority {
		logger.Info(fmt.Sprintf("- queue: %s, priority: %s", logger.Magenta.Add(q), logger.Green.Add(strconv.Itoa(p))))
	}
	// logger.Info("handle function: ")
	// for types, handler := range w.queue {
	// 	logger.Info(fmt.Sprintf(" task: %s ==> %s", types, getName(handler)))
	// }
	// fmt.Println()
	return w.server.Run(w.mux)
}
