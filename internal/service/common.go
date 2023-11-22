package service

import (
	"swclabs/swiftcart/internal/schema"
	"swclabs/swiftcart/internal/tasks"
	"swclabs/swiftcart/pkg/x/worker"
)

func Ping() {
	// log.Println("PONG")
}

func HealthCheck() schema.HealthCheckResponse {
	return schema.HealthCheckResponse{
		Status: "ok",
	}
}

func WorkerCheck() error {
	return worker.Exec(tasks.CriticalQueue, worker.NewTask(
		tasks.WorkerHealthCheck,
		1,
	))
}
