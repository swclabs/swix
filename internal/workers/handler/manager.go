// Package handler implements handler of worker
package handler

import (
	"context"
	"encoding/json"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/manager"

	"swclabs/swix/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

var _ = boot.Controller(NewManager)

// NewManager creates a new Manager object
func NewManager(handler manager.IManager) IManager {
	return &Manager{
		handler: handler,
	}
}

// IManager is an interface for Manager
type IManager interface {
	HandleSignUp() (string, worker.HandleFunc)
	HandleOAuth2SaveUser() (string, worker.HandleFunc)
	HandleUpdateUserInfo() (string, worker.HandleFunc)
}

// Manager struct define the Manager object
type Manager struct {
	manager.Task
	handler manager.IManager
}

// HandleSignUp handle sign up
func (manager *Manager) HandleSignUp() (string, worker.HandleFunc) {
	return worker.GetTaskName(manager.SignUp),
		func(_ context.Context, task *asynq.Task) error {
			var data dtos.SignUpRequest
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return manager.handler.SignUp(context.Background(), data)
		}
}

// HandleOAuth2SaveUser handle save user information from oauth2
func (manager *Manager) HandleOAuth2SaveUser() (string, worker.HandleFunc) {
	return worker.GetTaskName(manager.OAuth2SaveUser),
		func(_ context.Context, task *asynq.Task) error {
			var data dtos.OAuth2SaveUser
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return manager.handler.OAuth2SaveUser(context.Background(), data)
		}
}

// HandleUpdateUserInfo handle update user information
func (manager *Manager) HandleUpdateUserInfo() (string, worker.HandleFunc) {
	return worker.GetTaskName(manager.UpdateUserInfo),
		func(_ context.Context, task *asynq.Task) error {
			var data dtos.UserUpdate
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return manager.handler.UpdateUserInfo(context.Background(), data)
		}
}
