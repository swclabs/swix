package common

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

var _ ICommonService = (*Task)(nil)

type Task struct {
	worker worker.IWorkerClient
}

func (t *Task) CallTask() ICommonService {
	return t
}

func (t *Task) HealthCheck(ctx context.Context) domain.HealthCheckRes {
	return domain.HealthCheckRes{
		Status: "Ok",
	}
}

func (t *Task) WorkerCheck(ctx context.Context, num int64) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.WorkerCheck),
		1,
	))
}

func (t *Task) WorkerCheckResult(ctx context.Context, num int64) (string, error) {
	result, err := t.worker.ExecGetResult(ctx, queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.WorkerCheckResult),
		1,
	))
	if err != nil {
		return "", err
	}
	return string(result), err
}
