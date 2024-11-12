// Package authentication implements handler of worker
package authentication

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/service/authentication"
	"github.com/swclabs/swipex/pkg/lib/worker"
)

var _ = app.Controller(NewHandler)

// NewHandler creates a new Authentication object
func NewHandler(handler authentication.IAuthentication) *Handler {
	return &Handler{
		handler: handler,
	}
}

// Handler struct define the Handler object
type Handler struct {
	handler authentication.IAuthentication
}

// SignUp handle sign up
func (auth *Handler) SignUp(c worker.Context) error {
	var data dtos.SignUpRequest
	if err := json.Unmarshal(c.Payload(), &data); err != nil {
		return err
	}
	return auth.handler.SignUp(context.Background(), data)
}

// OAuth2SaveUser handle save user information from oauth2
func (auth *Handler) OAuth2SaveUser(c worker.Context) error {
	var data dtos.OAuth2SaveUser
	if err := json.Unmarshal(c.Payload(), &data); err != nil {
		return err
	}
	ID, err := auth.handler.OAuth2SaveUser(context.Background(), data)
	if err != nil {
		return err
	}
	return c.Return([]byte(strconv.FormatInt(ID, 10)))
}

// UpdateUserInfo handle update user information
func (auth *Handler) UpdateUserInfo(c worker.Context) error {
	var data dtos.UserUpdate
	if err := json.Unmarshal(c.Payload(), &data); err != nil {
		return err
	}
	return auth.handler.UpdateUserInfo(context.Background(), data)
}
