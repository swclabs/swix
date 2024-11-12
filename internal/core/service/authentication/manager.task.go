package authentication

import (
	"context"
	"mime/multipart"
	"strconv"

	"github.com/swclabs/swipex/internal/config"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/internal/workers/queue"
	"github.com/swclabs/swipex/pkg/lib/worker"
)

var _ IAuthentication = (*Task)(nil)

// Task struct for auth service
type Task struct {
	worker  worker.IWorkerClient
	service IAuthentication
}

// UseTask creates a new Task object wrapping the provided service
func UseTask(service IAuthentication) IAuthentication {
	return &Task{
		worker:  worker.NewClient(config.RedisHost, config.RedisPort, config.RedisPassword),
		service: service,
	}
}

// SignUp user to access system, return error if exist
func (t *Task) SignUp(ctx context.Context, req dtos.SignUpRequest) error {
	return t.worker.Exec(ctx,
		queue.DefaultQueue,
		worker.NewTask(
			"auth.SignUp",
			req,
		),
	)
}

// UpdateUserInfo update user information
func (t *Task) UpdateUserInfo(ctx context.Context, req dtos.UserUpdate) error {
	return t.worker.Exec(ctx,
		queue.DefaultQueue,
		worker.NewTask(
			"auth.UpdateUserInfo",
			req,
		),
	)
}

// OAuth2SaveUser save user information from oauth2
func (t *Task) OAuth2SaveUser(ctx context.Context, req dtos.OAuth2SaveUser) (int64, error) {
	result, err := t.worker.ExecGetResult(ctx,
		queue.DefaultQueue,
		worker.NewTask(
			"auth.OAuth2SaveUser",
			req,
		),
	)
	if err != nil {
		return -1, err
	}
	return strconv.ParseInt(string(result), 10, 64)

}

// Login user to access system, return token if success
func (t *Task) Login(ctx context.Context, req dtos.LoginRequest) (string, error) {
	return t.service.Login(ctx, req)
}

// CheckLoginEmail check if email is exist
func (t *Task) CheckLoginEmail(ctx context.Context, email string) error {
	return t.service.CheckLoginEmail(ctx, email)
}

// UserInfo get user information
func (t *Task) UserInfo(ctx context.Context, email string) (*model.Users, error) {
	return t.service.UserInfo(ctx, email)
}

// UploadAvatar upload avatar to database
func (t *Task) UploadAvatar(email string, fileHeader *multipart.FileHeader) error {
	return t.service.UploadAvatar(email, fileHeader)
}
