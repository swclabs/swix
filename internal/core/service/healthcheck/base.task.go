package healthcheck

import (
	"context"

	"github.com/swclabs/swipex/internal/config"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/workers/queue"
	"github.com/swclabs/swipex/pkg/lib/worker"
)

var _ IService = (*Task)(nil)

// Task struct for base service
type Task struct {
	worker  worker.IWorkerClient
	service IService
}

// UseTask creates a new Task object wrapping the provided service
func UseTask(service IService) IService {
	return &Task{
		worker:  worker.NewClient(config.RedisHost, config.RedisPort, config.RedisPassword),
		service: service,
	}
}

// HealthCheck implements IbaseService.
func (t *Task) HealthCheck(_ context.Context) dtos.HealthCheck {
	return dtos.HealthCheck{
		Status: "Ok",
	}
}

// WorkerCheck implements IbaseService.
func (t *Task) WorkerCheck(ctx context.Context, _ int64) error {
	return t.worker.Exec(ctx, queue.DefaultQueue, worker.NewTask(
		"healthcheck.WorkerCheck",
		1,
	))
}

// WorkerCheckResult implements IbaseService.
func (t *Task) WorkerCheckResult(ctx context.Context, num int64) (string, error) {
	result, err := t.worker.ExecGetResult(ctx, queue.DefaultQueue, worker.NewTask(
		"healthcheck.WorkerCheckResult",
		num,
	))
	if err != nil {
		return "", err
	}
	return string(result), err
}
