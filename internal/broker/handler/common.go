package handler

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/service"
	"github.com/swclabs/swipe-api/internal/tasks"
	"github.com/swclabs/swipe-api/pkg/tools/worker"
)

type CommonHandler struct {
	tasks.CommonTask                       // embedded delay function here
	handler          domain.ICommonService // create handler for services
}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{
		handler: service.NewCommonService(),
	}
}

func (common *CommonHandler) HandleHealthCheck() (taskName string, fn worker.HandleFunc) {
	// get task name from delay function
	taskName = worker.GetTaskName(common.DelayWorkerCheck)
	// implement handler function base on delay function
	return taskName, func(_ context.Context, task *asynq.Task) error {
		var num int64
		if err := json.Unmarshal(task.Payload(), &num); err != nil {
			return err
		}
		return common.handler.WorkerCheck(num)
	}
}
