package service

import (
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/tools/worker"
)

type IAccountManagement interface {
	DelaySignUp(req *domain.SignUpReq) error
	DelayUpdateUserInfo(req *domain.UserUpdate) error
	DelayOAuth2SaveUser(req *domain.OAuth2SaveUser) error
}

var _ IAccountManagement = (*AccountManagementTask)(nil)

type AccountManagementTask struct {
}

func NewAccountManagementTask() *AccountManagementTask {
	return &AccountManagementTask{}
}

func (t *AccountManagementTask) DelaySignUp(req *domain.SignUpReq) error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelaySignUp),
		req,
	))
}

func (t *AccountManagementTask) DelayUpdateUserInfo(req *domain.UserUpdate) error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayUpdateUserInfo),
		req,
	))
}

func (t *AccountManagementTask) DelayOAuth2SaveUser(req *domain.OAuth2SaveUser) error {
	return worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayOAuth2SaveUser),
		req,
	))
}
