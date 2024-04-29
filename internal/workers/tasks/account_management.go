package tasks

import (
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/tools/worker"
)

type IAccountManagement interface {
	DelaySignUp(req *domain.SignUpRequest) error
	DelayUpdateUserInfo(req *domain.UserUpdate) error
	DelayOAuth2SaveUser(req *domain.OAuth2SaveUser) error
}

var _ IAccountManagement = (*AccountManagement)(nil)

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
