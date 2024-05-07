package tasks

import (
	"context"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/tools/worker"
	"time"
)

type CommonTask struct {
}

func (common *CommonTask) DelayWorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
}

func (common *CommonTask) DelayWorkerCheckResult() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	result, err := worker.ExecGetResult(ctx, queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
	if err != nil {
		return "", err
	}
	return string(result), err
}
