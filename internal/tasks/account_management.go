package tasks

import (
	"context"
	"swclabs/swiftcart/internal/domain"
	"swclabs/swiftcart/pkg/worker"

	"github.com/hibiken/asynq"
)

// DEBUG: NOT DELETE THIS LINE
var _ IAccountManagement = NewAccountManagement()

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
		account.CallUpdateInfo:     account.WorkerUpdateInfo,
		account.CallOAuth2SaveUser: account.WorkerOAuth2SaveUser,
		account.CallNewUsers:       account.WorkerNewUsers,
	}
}

func (account *AccountManagement) UpdateUserInfo(req *domain.UserUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (account *AccountManagement) OAuth2SaveUser(req *domain.OAuth2SaveUser) error {
	//TODO implement me
	panic("implement me")
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
