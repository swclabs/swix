package service

import (
	"example/swiftcart/internal/schema"
	"example/swiftcart/internal/tasks"
	"example/swiftcart/pkg/lib/worker"
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
