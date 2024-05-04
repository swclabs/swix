package domain

import (
	"context"

	"gorm.io/gorm"
)

type ISuppliersRepository interface {
	Use(tx *gorm.DB) ISuppliersRepository

	Insert(ctx context.Context, sup Suppliers, addr Addresses) error
	InsertAddress(ctx context.Context, addr SuppliersAddress) error
	GetLimit(ctx context.Context, limit int) ([]Suppliers, error)
	GetByPhone(ctx context.Context, email string) (*Suppliers, error)
}
