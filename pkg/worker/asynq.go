package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/hibiken/asynq"
	"github.com/swclabs/swipe-api/pkg/logger"
)

type Priority map[string]int
type Queue map[string]func(context.Context, *asynq.Task) error
type Engine struct {
	server   *asynq.Server
	mux      *asynq.ServeMux
	priority Priority
	queue    Queue
}

var broker asynq.RedisClientOpt

func getName(input string) string {
	paths := strings.Split(input, "/")
	return paths[len(paths)-1]
}

func SetBroker(host, port, password string) {
	broker = asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", host, port), // Redis server address
		Password: password,                         // Redis password
	}
}

func NewServer(priorityQueue Priority) *Engine {
	// asynq.NewServer(broker, asynq.Config{
	// 	// Specify how many concurrent workers to use.
	// 	Concurrency: concurrency,
	// 	// Specify multiple queues with different priority.
	// 	Queues: Queue{},
	// })
	return &Engine{
		server:   nil,
		mux:      asynq.NewServeMux(),
		queue:    Queue{},
		priority: priorityQueue,
	}
}

func (w *Engine) Queue(taskName string, fn func(context.Context, *asynq.Task) error) {
	w.queue[taskName] = fn
}

func (w *Engine) handleFunctions() {
	for k, v := range w.queue {
		w.mux.HandleFunc(k, v)
	}
}

func (w *Engine) Run(concurrency int) error {
	// Create a new Asynq server
	w.server = asynq.NewServer(broker, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: concurrency,
		// Specify multiple queues with different priority.
		Queues: w.priority,
	})
	w.handleFunctions()

	logger.Banner("Launching a asynchronous worker with the following settings:")
	logger.Broker("redis", broker.Addr)
	for q, p := range w.priority {
		logger.Queue(q, p)
	}
	logger.Banner("Handle Function: ")
	for types, handler := range w.queue {
		logger.HandleFunc(types, getName(runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()))
	}
	fmt.Println()
	return w.server.Run(w.mux)
}

func Exec(queue string, task *asynq.Task) error {
	// Create a new Asynq client.
	client := asynq.NewClient(broker)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err.Error())
		}
	}(client)
	// Process the task immediately in critical queue.
	_, err := client.Enqueue(
		task,               // task payload
		asynq.Queue(queue), // set queue for task
	)
	if err != nil {
		return err
	}
	return nil

}

func NewTask(typename string, data interface{}) *asynq.Task {
	payload, _ := json.Marshal(data)
	return asynq.NewTask(typename, payload)
}

func Delay(delay *time.Duration, queue string, task *asynq.Task) error {
	// Create a new Asynq client.
	client := asynq.NewClient(broker)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err.Error())
		}
	}(client)
	if _, err := client.Enqueue(
		task,                    // task payload
		asynq.Queue(queue),      // set queue for task
		asynq.ProcessIn(*delay), // set time to process task
	); err != nil {
		return err
	}
	return nil
}
