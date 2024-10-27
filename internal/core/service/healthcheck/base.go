// Package base implements base services
package healthcheck

import (
	"context"
	"fmt"
	"strconv"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
)

var _ IService = (*Service)(nil)

// Service struct for base service

// New creates a new Service object
func New() IService {
	return &Service{}
}

var _ = app.Service(New)

type Service struct{}

// WorkerCheckResult implements IbaseService.
func (base *Service) WorkerCheckResult(ctx context.Context, num int64) (string, error) {
	return strconv.Itoa(int(num)), base.WorkerCheck(ctx, num)
}

// HealthCheck implements IbaseService.
func (base *Service) HealthCheck(_ context.Context) dtos.HealthCheck {
	return dtos.HealthCheck{
		Status: "ok",
	}
}

// WorkerCheck implements IbaseService.
func (base *Service) WorkerCheck(_ context.Context, num int64) error {
	fmt.Printf("HealthCheck Number: %d\n", int(num))
	return nil
}
