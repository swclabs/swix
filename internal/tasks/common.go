package tasks

import (
	"swclabs/swiftcart/internal/messaging/queue"
	"swclabs/swiftcart/pkg/worker"
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
