package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"reflect"
	"runtime"
	"strings"
	"time"
)

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
