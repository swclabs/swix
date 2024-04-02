package service

import (
	"context"
	"fmt"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/helper/tasks"
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
