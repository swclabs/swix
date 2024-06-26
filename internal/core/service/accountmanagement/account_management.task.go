package accountmanagement

import (
	"context"
	"fmt"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

var _ IAccountManagement = (*Task)(nil)

type Task struct {
	worker worker.IWorkerClient
}

func (t *Task) CallTask() IAccountManagement {
	return t
}

func (t *Task) SignUp(_ context.Context, req domain.SignUpReq) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.SignUp),
		req,
	))
}

func (t *Task) UpdateUserInfo(_ context.Context, req domain.UserUpdate) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.UpdateUserInfo),
		req,
	))
}

func (t *Task) OAuth2SaveUser(_ context.Context, req domain.OAuth2SaveUser) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.OAuth2SaveUser),
		req,
	))
}

func (t *Task) Login(ctx context.Context, req domain.LoginReq) (string, error) {
	return "", fmt.Errorf("%s: function not available", worker.GetTaskName(t.Login))
}

func (t *Task) CheckLoginEmail(ctx context.Context, email string) error {
	return fmt.Errorf(
		"%s: function not available", worker.GetTaskName(t.CheckLoginEmail))
}

func (t *Task) UserInfo(ctx context.Context, email string) (*domain.UserInfo, error) {
	return nil, fmt.Errorf(
		"%s: function not available", worker.GetTaskName(t.UserInfo))
}

func (t *Task) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	return fmt.Errorf(
		"%s: function not available", worker.GetTaskName(t.UploadAvatar))
}

func (t *Task) UploadAddress(ctx context.Context, data domain.Addresses) error {
	return fmt.Errorf(
		"%s: function not available", worker.GetTaskName(t.UploadAddress))
}
