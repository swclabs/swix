package domain

import (
	"context"

	"gorm.io/gorm"
)

type IAccountRepository interface {
	Use(tx *gorm.DB) IAccountRepository
	GetByEmail(ctx context.Context, email string) (*Account, error)
	Insert(ctx context.Context, acc *Account) error
	SaveInfo(ctx context.Context, acc *Account) error
}
