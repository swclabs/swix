package common

import (
	"context"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

type Task struct {
	worker worker.IWorkerClient
}

func (t *Task) DelayWorkerCheck() error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayWorkerCheck),
		1,
	))
}

func (t *Task) DelayWorkerCheckResult(ctx context.Context) (string, error) {
	result, err := t.worker.ExecGetResult(ctx, queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayWorkerCheck),
		1,
	))
	if err != nil {
		return "", err
	}
	return string(result), err
}
