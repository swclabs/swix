package domain

import (
	"context"

	"gorm.io/gorm"
)

// Addresses Table
type Addresses struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Uuid     string `json:"uuid" gorm:"column:uuid"`
	City     string `json:"city" gorm:"column:city"`
	Ward     string `json:"ward" gorm:"column:ward"`
	District string `json:"district" gorm:"column:district"`
	Street   string `json:"street" gorm:"column:street"`
}

type IAddressRepository interface {
	Use(tx *gorm.DB) IAddressRepository
	Insert(ctx context.Context, data *Addresses) error
}
