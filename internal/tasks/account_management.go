package tasks

import (
	"context"
	"swclabs/swiftcart/pkg/worker"

	"github.com/hibiken/asynq"
)

type AccountManagement struct {
	CallUploadImage    string
	CallUpdateInfo     string
	CallOAuth2SaveUser string
	CallNewUsers       string
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		CallUploadImage:    worker.NewTypename("UploadImage"),
		CallUpdateInfo:     worker.NewTypename("UpdateInfo"),
		CallOAuth2SaveUser: worker.NewTypename("OAuth2SaveUser"),
		CallNewUsers:       worker.NewTypename("NewUsers"),
	}
}

func AccountManagementPath() worker.Path {
	account := NewAccountManagement()
	return worker.Path{
		account.CallUploadImage:    account.UploadImage,
		account.CallUpdateInfo:     account.UpdateInfo,
		account.CallOAuth2SaveUser: account.OAuth2SaveUser,
		account.CallNewUsers:       account.NewUsers,
	}
}

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
