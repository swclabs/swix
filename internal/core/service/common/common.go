package common

import (
	"context"
	"fmt"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/lib/worker"
)

type CommonService struct {
	Task
}

func New(
	client *worker.Client,
) *CommonService {
	return &CommonService{
		Task{
			worker: client,
		},
	}
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
