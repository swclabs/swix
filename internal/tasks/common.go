package tasks

import (
	"github.com/swclabs/swipe-server/internal/broker/queue"
	"github.com/swclabs/swipe-server/pkg/tools/worker"
)

type CommonTask struct {
}

func (common *CommonTask) DelayWorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(common.DelayWorkerCheck),
		1,
	))
}
