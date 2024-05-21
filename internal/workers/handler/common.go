package handler

import (
	"context"
	"encoding/json"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

type CommonHandler struct {
	service.CommonTask                       // embedded delay function here
	handler            domain.ICommonService // create handler for services
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
		if err := common.handler.WorkerCheck(context.Background(), num); err != nil {
			return err
		}
		_, err := task.ResultWriter().Write([]byte("HandleHealthCheck: success"))
		return err
	}
}
