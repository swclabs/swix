package common

import (
	"context"
	"fmt"
	"swclabs/swipecore/internal/core/domain"
)

type CommonService struct {
	Task
}

func New() *CommonService {
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
