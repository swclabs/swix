package tasks

import (
	"github.com/swclabs/swipe-api/internal/domain"
)

type AccountManagement struct {
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{}
}

func (t *AccountManagement) DelaySignUp(req *domain.SignUpRequest) error {
	return nil
}
