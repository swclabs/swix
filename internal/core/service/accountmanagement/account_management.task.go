package accountmanagement

import (
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

var _ IAccountManagementTask = (*Task)(nil)

type Task struct {
	worker worker.IWorkerClient
}

func (t *Task) DelaySignUp(req domain.SignUpReq) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelaySignUp),
		req,
	))
}

func (t *Task) DelayUpdateUserInfo(req domain.UserUpdate) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayUpdateUserInfo),
		req,
	))
}

func (t *Task) DelayOAuth2SaveUser(req domain.OAuth2SaveUser) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.DelayOAuth2SaveUser),
		req,
	))
}
