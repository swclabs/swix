package domain

import (
	"context"

	"gorm.io/gorm"
)

const SuppliersTable = "suppliers"
const SuppliersAddressTable = "supplier_address"

// Suppliers table
type Suppliers struct {
	Id          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Email       string `json:"email" gorm:"column:email"`
}

// SuppliersAddress suppliers address table
// CREATE TABLE "supplier_address" (
//
//		"supplier_id" bigint PRIMARY KEY,
//		"address_uuid" bigint NOT NULL
//	  );
type SuppliersAddress struct {
	SuppliersID string `json:"suppliers_id" gorm:"column:suppliers_id"`
	AddressUuiD string `json:"address_uuid" gorm:"column:address_uuid"`
}

type SuppliersListResponse struct {
	Data []Suppliers `json:"data"`
}

type SuppliersRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"number,required"`
	Email       string `json:"email" validate:"email,required"`
	City        string `json:"city" gorm:"column:city"`
	Ward        string `json:"ward" gorm:"column:ward"`
	District    string `json:"district" gorm:"column:district"`
	Street      string `json:"street" gorm:"column:street"`
}

type ISuppliersRepository interface {
	Use(tx *gorm.DB) ISuppliersRepository

	Insert(ctx context.Context, sup Suppliers, addr Addresses) error
	InsertAddress(ctx context.Context, addr SuppliersAddress) error
	GetLimit(ctx context.Context, limit int) ([]Suppliers, error)
	GetByPhone(ctx context.Context, email string) (*Suppliers, error)
}
