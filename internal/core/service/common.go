package service

import (
	"context"
	"fmt"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/tasks"
)

type CommonService struct {
	tasks.CommonTask
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (common *CommonService) HealthCheck(ctx context.Context) domain.HealthCheckResponse {
	return domain.HealthCheckResponse{
		Status: "ok",
	}
}

func (common *CommonService) WorkerCheck(ctx context.Context, num int64) error {
	fmt.Printf("HealthCheck Number: %d\n", int(num))
	return nil
}
