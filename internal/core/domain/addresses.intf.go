package domain

import (
	"context"

	"gorm.io/gorm"
)

type IAddressRepository interface {
	Use(tx *gorm.DB) IAddressRepository
	Insert(ctx context.Context, data *Addresses) error
}
