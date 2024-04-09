package tasks

import (
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/workers/queue"
	"swclabs/swipe-api/pkg/tools/worker"
)

type AccountManagement struct {
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{}
}

func (t *AccountManagement) DelaySignUp(req *domain.SignUpRequest) error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelaySignUp),
		req,
	))
}

func (t *AccountManagement) DelayUpdateUserInfo(req *domain.UserUpdate) error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayUpdateUserInfo),
		req,
	))
}

func (t *AccountManagement) DelayOAuth2SaveUser(req *domain.OAuth2SaveUser) error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayOAuth2SaveUser),
		req,
	))
}
