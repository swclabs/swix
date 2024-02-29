package service

import (
	"fmt"

	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/tasks"
)

type CommonService struct {
	tasks.CommonTask
}

func NewCommonService() *CommonService {
	return &CommonService{}
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
