package tasks

import (
	"github.com/swclabs/swipe-api/internal/broker/queue"
	"github.com/swclabs/swipe-api/pkg/worker"
)

type CommonTask struct {
}

func NewCommonTask() *CommonTask {
	return &CommonTask{}
}

func (common *CommonTask) WorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		queue.WorkerHealthCheck,
		1,
	))
}
