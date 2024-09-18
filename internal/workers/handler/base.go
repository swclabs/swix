// Package handler implements handler of worker
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/service/base"

	"swclabs/swix/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

var _ IBaseHandler = (*BaseHandler)(nil)
var _ = app.Controller(NewBase)

// NewBase creates a new base object
func NewBase(_base base.IService) IBaseHandler {
	return &BaseHandler{
		handler: _base,
	}
}

// IBaseHandler is an interface for Base.
type IBaseHandler interface {
	HandleHealthCheck() (taskName string, fn worker.HandleFunc)
}

// BaseHandler struct define the base object
type BaseHandler struct {
	base.Task               // embedded delay function here
	handler   base.IService // create handler for services
}

// HandleHealthCheck handle health check
func (base *BaseHandler) HandleHealthCheck() (taskName string, fn worker.HandleFunc) {
	// get task name from delay function
	taskName = worker.GetTaskName(base.WorkerCheckResult)
	// implement handler function base on delay function
	return taskName, func(_ context.Context, task *asynq.Task) error {
		var num int64
		if err := json.Unmarshal(task.Payload(), &num); err != nil {
			return err
		}
		result, err := base.handler.WorkerCheckResult(context.Background(), num)
		if err != nil {
			return err
		}
		_, err = task.ResultWriter().Write(
			[]byte(fmt.Sprintf("HandleHealthCheck with param '%s': success", result)))
		return err
	}
}
