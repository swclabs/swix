package common

import (
	"context"
	"fmt"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/lib/worker"
)

var _ ICommonService = (*Service)(nil)

type Service struct {
	Task *Task
}

func New(
	client worker.IWorkerClient,
) ICommonService {
	return &Service{
		&Task{
			worker: client,
		},
	}
}

func (common *Service) CallTask() ICommonService {
	return common.Task
}

func (common *Service) WorkerCheckResult(ctx context.Context, num int64) (string, error) {
	return strconv.Itoa(int(num)), common.WorkerCheck(ctx, num)
}

func (common *Service) HealthCheck(ctx context.Context) domain.HealthCheckRes {
	return domain.HealthCheckRes{
		Status: "ok",
	}
}

func (common *Service) WorkerCheck(ctx context.Context, num int64) error {
	fmt.Printf("HealthCheck Number: %d\n", int(num))
	return nil
}
