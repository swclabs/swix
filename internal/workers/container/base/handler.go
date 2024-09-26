// Package handler implements handler of worker
package base

import (
	"context"
	"encoding/json"
	"fmt"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/service/base"

	"github.com/hibiken/asynq"
)

var _ = app.Controller(NewHandler)

// NewHandler creates a new base object
func NewHandler(_base base.IService) *Handler {
	return &Handler{
		handler: _base,
	}
}

// Handler struct define the base object
type Handler struct {
	handler base.IService // create handler for services
}

// HandleHealthCheck handle health check
func (base *Handler) WorkerCheckResult(_ context.Context, task *asynq.Task) error {
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
