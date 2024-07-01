package common

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ICommonService : Common utility methods for the service.
// Actor: System
type ICommonService interface {
	// HealthCheck performs a health check on the service.
	// ctx is the context to manage the request's lifecycle.
	// Returns a HealthCheck object with the health check status.
	HealthCheck(ctx context.Context) domain.HealthCheck

	// WorkerCheck checks the status of a worker.
	// ctx is the context to manage the request's lifecycle.
	// num specifies the worker number to check.
	// Returns an error if any issues occur during the check process.
	WorkerCheck(ctx context.Context, num int64) error

	WorkerCheckResult(ctx context.Context, num int64) (string, error)
}
