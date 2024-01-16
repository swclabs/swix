package handler

import (
	"context"
	"encoding/json"
	"swclabs/swiftcart/internal/service"

	"github.com/hibiken/asynq"
)

type CommonHandler struct {
	handler *service.CommonService
}

func NewCommonHandler() *CommonHandler {
	return &CommonHandler{
		handler: service.NewCommonService(),
	}
}

func (common *CommonHandler) HandleHealthCheck(_ context.Context, task *asynq.Task) error {
	var num int64
	if err := json.Unmarshal(task.Payload(), &num); err != nil {
		return err
	}
	return common.handler.WorkerCheck(num)
}
