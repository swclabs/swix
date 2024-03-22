package tasks

import (
	"github.com/swclabs/swipe-api/internal/workers/queue"
	"github.com/swclabs/swipe-api/pkg/tools/worker"
)

type CommonTask struct {
}

func (common *CommonTask) DelayWorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
}
