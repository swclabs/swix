package tasks

import (
	"context"
	"github.com/hibiken/asynq"
	"swclabs/swiftcart/internal/domain"
)

type IAccountManagement interface {
	// base handle function
	UpdateUserInfo(req *domain.UserUpdate) error
	OAuth2SaveUser(req *domain.OAuth2SaveUser) error

	// worker handle function
	WorkerUpdateInfo(ctx context.Context, task *asynq.Task) error
	WorkerOAuth2SaveUser(ctx context.Context, task *asynq.Task) error
	WorkerNewUsers(ctx context.Context, task *asynq.Task) error
}
