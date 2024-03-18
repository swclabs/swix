package tasks

import (
	"github.com/swclabs/swipe-server/internal/broker/queue"
	"github.com/swclabs/swipe-server/internal/domain"
	"github.com/swclabs/swipe-server/pkg/tools/worker"
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
