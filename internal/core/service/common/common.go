package common

import (
	"context"
	"fmt"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
)

var _ ICommonService = (*Service)(nil)

type Service struct{}

func New() ICommonService {
	return &Service{}
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
