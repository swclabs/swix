package tasks

import (
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/tools/worker"
)

type CommonTask struct {
}

func (common *CommonTask) DelayWorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
}
