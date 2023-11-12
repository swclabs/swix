package service

import (
	"example/komposervice/internal/schema"
	"example/komposervice/internal/tasks"
	"example/komposervice/pkg/lib/worker"
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
