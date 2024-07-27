// Package common implements common services
package common

import (
	"context"
	"fmt"
	"strconv"
	"swclabs/swipecore/internal/core/domain/dtos"
)

var _ ICommonService = (*Service)(nil)

// Service struct for common service
type Service struct{}

// New creates a new Service object
func New() ICommonService {
	return &Service{}
}

// WorkerCheckResult implements ICommonService.
func (common *Service) WorkerCheckResult(ctx context.Context, num int64) (string, error) {
	return strconv.Itoa(int(num)), common.WorkerCheck(ctx, num)
}

// HealthCheck implements ICommonService.
func (common *Service) HealthCheck(_ context.Context) dtos.HealthCheck {
	return dtos.HealthCheck{
		Status: "ok",
	}
}

// WorkerCheck implements ICommonService.
func (common *Service) WorkerCheck(_ context.Context, num int64) error {
	fmt.Printf("HealthCheck Number: %d\n", int(num))
	return nil
}
