package accountmanagement

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/pkg/lib/worker"
)

var _ IAccountManagement = (*Task)(nil)

type Task struct {
	worker  worker.IWorkerClient
	service IAccountManagement
}

func QueueOf(service IAccountManagement) IAccountManagement {
	return &Task{
		worker:  worker.NewClient(config.LoadEnv()),
		service: service,
	}
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
	return t.service.Login(ctx, req)
}

func (t *Task) CheckLoginEmail(ctx context.Context, email string) error {
	return t.service.CheckLoginEmail(ctx, email)
}

func (t *Task) UserInfo(ctx context.Context, email string) (*domain.UserInfo, error) {
	return t.service.UserInfo(ctx, email)
}

func (t *Task) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	return t.service.UploadAvatar(email, fileHeader)
}

func (t *Task) UploadAddress(ctx context.Context, data domain.Addresses) error {
	return t.service.UploadAddress(ctx, data)
}
