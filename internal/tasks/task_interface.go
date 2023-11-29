package tasks

import (
	"context"

	"github.com/hibiken/asynq"
)

type IAccountManagement interface {
	UploadImage(ctx context.Context, task *asynq.Task) error
	UpdateInfo(ctx context.Context, task *asynq.Task) error
	OAuth2SaveUser(ctx context.Context, task *asynq.Task) error
	NewUsers(ctx context.Context, task *asynq.Task) error
}
