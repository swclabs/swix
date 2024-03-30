package tasks

import (
	"swclabs/swipe-api/internal/workers/queue"
	"swclabs/swipe-api/pkg/tools/worker"
)

type CommonTask struct {
}

func (common *CommonTask) DelayWorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
}
