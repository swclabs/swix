package domain

import (
	"context"

	"gorm.io/gorm"
)

// Account table
type Account struct {
	Username  string `json:"username" gorm:"column:username"`
	Role      string `json:"role" gorm:"column:role"`
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password" gorm:"column:password"`
	CreatedAt string `json:"created_at" gorm:"column:created"`
	Type      string `json:"type" gorm:"column:type"`
}

type IAccountRepository interface {
	Use(tx *gorm.DB) IAccountRepository
	GetByEmail(ctx context.Context, email string) (*Account, error)
	Insert(ctx context.Context, acc *Account) error
	SaveInfo(ctx context.Context, acc *Account) error
}
