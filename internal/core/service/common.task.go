package service

import (
	"context"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

type CommonTask struct {
}

func (common *CommonTask) DelayWorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
}

func (common *CommonTask) DelayWorkerCheckResult(ctx context.Context) (string, error) {
	result, err := worker.ExecGetResult(ctx, queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
	if err != nil {
		return "", err
	}
	return string(result), err
}
