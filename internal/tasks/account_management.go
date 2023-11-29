package tasks

import (
	"context"

	"github.com/hibiken/asynq"
)

type AccountManagement struct{}

func (account *AccountManagement) UploadImage(ctx context.Context, task *asynq.Task) error {
	//TODO implement me
	panic("implement me")
}

func (account *AccountManagement) UpdateInfo(ctx context.Context, task *asynq.Task) error {
	//TODO implement me
	panic("implement me")
}

func (account *AccountManagement) OAuth2SaveUser(ctx context.Context, task *asynq.Task) error {
	//TODO implement me
	panic("implement me")
}

func (account *AccountManagement) NewUsers(ctx context.Context, task *asynq.Task) error {
	//TODO implement me
	panic("implement me")
}
