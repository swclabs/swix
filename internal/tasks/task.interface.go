package tasks

import (
	"context"
	"swclabs/swiftcart/internal/schema"

	"github.com/hibiken/asynq"
)

type IAccountManagement interface {
	// base handle function
	UpdateUserInfo(req *schema.UserUpdate) error
	OAuth2SaveUser(req *schema.OAuth2SaveUser) error

	// worker handle function
	WorkerUpdateInfo(ctx context.Context, task *asynq.Task) error
	WorkerOAuth2SaveUser(ctx context.Context, task *asynq.Task) error
	WorkerNewUsers(ctx context.Context, task *asynq.Task) error
}
