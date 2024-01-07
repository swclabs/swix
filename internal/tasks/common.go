package tasks

import (
	"swclabs/swiftcart/internal/delivery/msg/queue"
	"swclabs/swiftcart/pkg/worker"
)

type CommonTask struct {
	WorkerHealthCheck string
}

func NewCommonTask() *CommonTask {
	return &CommonTask{
		WorkerHealthCheck: "Worker#HealthCheck",
	}
}

func (common *CommonTask) WorkerCheck() error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		common.WorkerHealthCheck,
		1,
	))
}
