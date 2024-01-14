package handler

import (
	"context"
	"encoding/json"
	"swclabs/swiftcart/internal/service"
	"swclabs/swiftcart/internal/tasks"

	"github.com/hibiken/asynq"
)

var common = NewCommonHandler()

type CommonHandler struct {
	taskName *tasks.CommonTask
	handler  *service.CommonService
}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{
		taskName: tasks.NewCommonTask(),
		handler:  service.NewCommonService(),
	}
}

func (common *CommonHandler) HandleHealthCheck(_ context.Context, task *asynq.Task) error {
	var num int64
	if err := json.Unmarshal(task.Payload(), &num); err != nil {
		return err
	}
	return common.handler.WorkerCheck(num)
}
