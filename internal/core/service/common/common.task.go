package common

import (
	"context"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

var _ ICommonService = (*Task)(nil)

type Task struct {
	worker  worker.IWorkerClient
	service ICommonService
}

func UseTask(service ICommonService) ICommonService {
	return &Task{
		worker:  worker.NewClient(config.RedisHost, config.RedisPort, config.RedisPassword),
		service: service,
	}
}

func (t *Task) HealthCheck(_ context.Context) domain.HealthCheck {
	return domain.HealthCheck{
		Status: "Ok",
	}
}

func (t *Task) WorkerCheck(_ context.Context, _ int64) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.WorkerCheck),
		1,
	))
}

func (t *Task) WorkerCheckResult(ctx context.Context, num int64) (string, error) {
	result, err := t.worker.ExecGetResult(ctx, queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.WorkerCheckResult),
		num,
	))
	if err != nil {
		return "", err
	}
	return string(result), err
}
