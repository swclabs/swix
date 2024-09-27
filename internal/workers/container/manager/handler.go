// Package manager implements handler of worker
package manager

import (
	"context"
	"encoding/json"
	"strconv"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/manager"
	"swclabs/swix/pkg/lib/worker"
)

var _ = app.Controller(NewHandler)

// NewHandler creates a new Manager object
func NewHandler(handler manager.IManager) *Handler {
	return &Handler{
		handler: handler,
	}
}

// Handler struct define the Handler object
type Handler struct {
	handler manager.IManager
}

// SignUp handle sign up
func (manager *Handler) SignUp(c worker.Context) error {
	var data dtos.SignUpRequest
	if err := json.Unmarshal(c.Payload(), &data); err != nil {
		return err
	}
	return manager.handler.SignUp(context.Background(), data)
}

// OAuth2SaveUser handle save user information from oauth2
func (manager *Handler) OAuth2SaveUser(c worker.Context) error {
	var data dtos.OAuth2SaveUser
	if err := json.Unmarshal(c.Payload(), &data); err != nil {
		return err
	}
	ID, err := manager.handler.OAuth2SaveUser(context.Background(), data)
	if err != nil {
		return err
	}
	return c.Return([]byte(strconv.FormatInt(ID, 10)))
}

// UpdateUserInfo handle update user information
func (manager *Handler) UpdateUserInfo(c worker.Context) error {
	var data dtos.UserUpdate
	if err := json.Unmarshal(c.Payload(), &data); err != nil {
		return err
	}
	return manager.handler.UpdateUserInfo(context.Background(), data)
}
