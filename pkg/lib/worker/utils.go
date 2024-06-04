package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var broker asynq.RedisClientOpt

// getName returns the name of the function
func getName(input interface{}) string {
	str := runtime.FuncForPC(reflect.ValueOf(input).Pointer()).Name()
	paths := strings.Split(str, "/")
	return paths[len(paths)-1]
}

// wait for task to complete and get result
func await(ctx context.Context, i *asynq.Inspector, queue, taskID string) (*asynq.TaskInfo, error) {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			taskInfo, err := i.GetTaskInfo(queue, taskID)
			if err != nil {
				return nil, err
			}
			if taskInfo.CompletedAt.IsZero() {
				continue
			}
			return taskInfo, nil
		case <-ctx.Done():
			taskInfo, err := i.GetTaskInfo(queue, taskID)
			if err != nil {
				return nil, err
			}
			return taskInfo, i.DeleteTask(queue, taskID)
		}
	}
}

// SetBroker set redis host and port to asynq.RedisClientOpt
func SetBroker(host, port, password string) {
	broker = asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", host, port), // Redis server address
		Password: password,                         // Redis password
	}
}

// GetTaskName get task name from function
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

// ExecGetResult executes tasks in the given queue
func ExecGetResult(ctx context.Context, queue string, task *asynq.Task) ([]byte, error) {
	// Create a new Asynq client.
	client := asynq.NewClient(broker)
	defer func(client *asynq.Client) {
		err := client.Close()
		if err != nil {
			panic(err.Error())
		}
	}(client)

	asyncResult, err := client.Enqueue(
		task,               // task payload
		asynq.Queue(queue), // set queue for task
		asynq.Retention(time.Duration(time.Second*30)), // store tasks when finished
	)
	if err != nil {
		return nil, err
	}

	inspector := asynq.NewInspector(broker)
	defer func(client *asynq.Inspector) {
		err := inspector.Close()
		if err != nil {
			panic(err.Error())
		}
	}(inspector)

	result, err := await(ctx, inspector, queue, asyncResult.ID)
	if err != nil {
		return nil, err
	}

	return result.Result, nil
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
