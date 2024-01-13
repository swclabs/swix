package handler

import (
	"context"
	"swclabs/swiftcart/pkg/worker"

	"github.com/hibiken/asynq"
)

type AccountManagement struct {
}

func AccountManagementPath() worker.Path {
	return worker.Path{}
}

func (account *AccountManagement) WorkerUpdateInfo(ctx context.Context, task *asynq.Task) error {
	//TODO implement me
	panic("implement me")
}

func (account *AccountManagement) WorkerOAuth2SaveUser(ctx context.Context, task *asynq.Task) error {
	//TODO implement me
	panic("implement me")
}

func (account *AccountManagement) WorkerNewUsers(ctx context.Context, task *asynq.Task) error {
	//TODO implement me
	panic("implement me")
}
