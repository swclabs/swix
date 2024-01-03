package service

import (
	"swclabs/swiftcart/internal/domain"
	"swclabs/swiftcart/internal/tasks"
	"swclabs/swiftcart/pkg/worker"
)

func Ping() {

}

func HealthCheck() domain.HealthCheckResponse {
	return domain.HealthCheckResponse{
		Status: "ok",
	}
}

func WorkerCheck() error {
	return worker.Exec(tasks.CriticalQueue, worker.NewTask(
		tasks.WorkerHealthCheck,
		1,
	))
}
