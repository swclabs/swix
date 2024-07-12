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

// Task struct for account management service
type Task struct {
	worker  worker.IWorkerClient
	service IAccountManagement
}

// UseTask creates a new Task object wrapping the provided service
func UseTask(service IAccountManagement) IAccountManagement {
	return &Task{
		worker:  worker.NewClient(config.RedisHost, config.RedisPort, config.RedisPassword),
		service: service,
	}
}

// SignUp user to access system, return error if exist
func (t *Task) SignUp(_ context.Context, req domain.SignUpSchema) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.SignUp),
		req,
	))
}

// UpdateUserInfo update user information
func (t *Task) UpdateUserInfo(_ context.Context, req domain.UserUpdate) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.UpdateUserInfo),
		req,
	))
}

// OAuth2SaveUser save user information from oauth2
func (t *Task) OAuth2SaveUser(_ context.Context, req domain.OAuth2SaveUser) error {
	return t.worker.Exec(queue.CriticalQueue, worker.NewTask(
		worker.GetTaskName(t.OAuth2SaveUser),
		req,
	))
}

// Login user to access system, return token if success
func (t *Task) Login(ctx context.Context, req domain.LoginSchema) (string, error) {
	return t.service.Login(ctx, req)
}

// CheckLoginEmail check if email is exist
func (t *Task) CheckLoginEmail(ctx context.Context, email string) error {
	return t.service.CheckLoginEmail(ctx, email)
}

// UserInfo get user information
func (t *Task) UserInfo(ctx context.Context, email string) (*domain.UserSchema, error) {
	return t.service.UserInfo(ctx, email)
}

// UploadAvatar upload avatar to database
func (t *Task) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	return t.service.UploadAvatar(email, fileHeader)
}

// UploadAddress upload address to database
func (t *Task) UploadAddress(ctx context.Context, data domain.Addresses) error {
	return t.service.UploadAddress(ctx, data)
}
