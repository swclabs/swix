package common

import (
	"context"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

type Task struct {
}

func (common *Task) DelayWorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
}

func (common *Task) DelayWorkerCheckResult(ctx context.Context) (string, error) {
	result, err := worker.ExecGetResult(ctx, queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
	if err != nil {
		return "", err
	}
	return string(result), err
}
