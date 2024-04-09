package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"swclabs/swipe-api/pkg/tools/logger"

	"github.com/hibiken/asynq"
)

// Priority is a queue priority define in asynq
type Priority map[string]int

// Queue is a map of queue name and handler function
type Queue map[string]func(context.Context, *asynq.Task) error

// HandleFunc is a sort of handler function type
type HandleFunc func(ctx context.Context, task *asynq.Task) error

// Engine includes all components of the asynq server
type Engine struct {
	server   *asynq.Server
	mux      *asynq.ServeMux
	priority Priority
	queue    Queue
}

var broker asynq.RedisClientOpt

// NewServer creates a new instance of the Worker consume server
func NewServer(priorityQueue Priority) *Engine {
	return &Engine{
		server:   nil,
		mux:      asynq.NewServeMux(),
		queue:    Queue{},
		priority: priorityQueue,
	}
}

// handleFunctions run all functions in the given path
func (w *Engine) handleFunctions() {
	for k, v := range w.queue {
		w.mux.HandleFunc(k, v)
	}
}

func (w *Engine) Queue(hfn func() (taskName string, fn HandleFunc)) {
	taskName, fn := hfn()
	w.queue[taskName] = fn
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
		logger.HandleFunc(types, getName(handler))
	}
	fmt.Println()
	return w.server.Run(w.mux)
}

// getName returns the name of the function
func getName(input interface{}) string {
	str := runtime.FuncForPC(reflect.ValueOf(input).Pointer()).Name()
	paths := strings.Split(str, "/")
	return paths[len(paths)-1]
}

// SetBroker set redis host and port to asynq.RedisClientOpt
func SetBroker(host, port, password string) {
	broker = asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", host, port), // Redis server address
		Password: password,                         // Redis password
	}
}

// GetTaskName get taskname from function
func GetTaskName(input interface{}) string {
	return getName(input)
}

// NewTask creates a new asynq.Task instance
// to be executed by worker consume handler.
func NewTask(taskName string, data interface{}) *asynq.Task {
	payload, _ := json.Marshal(data)
	return asynq.NewTask(taskName, payload)
}

// Exec executes tasks in the given queue
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

// Delay executes tasks after period of time.
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
