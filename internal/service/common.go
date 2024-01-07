package service

import (
	"fmt"
	"swclabs/swiftcart/internal/domain"
	"swclabs/swiftcart/internal/tasks"
)

func Ping() {

}

type CommonService struct {
	Task *tasks.CommonTask
}

func NewCommonService() *CommonService {
	return &CommonService{
		Task: tasks.NewCommonTask(),
	}
}

func (common *CommonService) HealthCheck() domain.HealthCheckResponse {
	return domain.HealthCheckResponse{
		Status: "ok",
	}
}

func (common *CommonService) WorkerCheck(num int64) error {
	fmt.Printf("HealthCheck Number: %d\n", int(num))
	return nil
}
