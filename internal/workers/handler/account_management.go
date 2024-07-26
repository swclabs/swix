// Package handler implements handler of worker
package handler

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain/dto"
	"swclabs/swipecore/internal/core/service/accountmanagement"

	"swclabs/swipecore/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

// AccountManagement struct define the AccountManagement object
type AccountManagement struct {
	accountmanagement.Task
	handler accountmanagement.IAccountManagement
}

// NewAccountManagementConsume creates a new AccountManagement object
func NewAccountManagementConsume(handler accountmanagement.IAccountManagement) *AccountManagement {
	return &AccountManagement{
		handler: handler,
	}
}

// HandleSignUp handle sign up
func (account *AccountManagement) HandleSignUp() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.SignUp),
		func(_ context.Context, task *asynq.Task) error {
			var data dto.SignUpRequest
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.SignUp(context.Background(), data)
		}
}

// HandleOAuth2SaveUser handle save user information from oauth2
func (account *AccountManagement) HandleOAuth2SaveUser() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.OAuth2SaveUser),
		func(_ context.Context, task *asynq.Task) error {
			var data dto.OAuth2SaveUser
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.OAuth2SaveUser(context.Background(), data)
		}
}

// HandleUpdateUserInfo handle update user information
func (account *AccountManagement) HandleUpdateUserInfo() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.UpdateUserInfo),
		func(_ context.Context, task *asynq.Task) error {
			var data dto.User
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.UpdateUserInfo(context.Background(), data)
		}
}
