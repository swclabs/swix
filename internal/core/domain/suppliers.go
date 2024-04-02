package domain

import "context"

// Suppliers table
type Suppliers struct {
	Id          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Email       string `json:"email" gorm:"column:email"`
}

// SuppliersAddress suppliers address table
type SuppliersAddress struct {
	SuppliersID string `json:"suppliers_id" gorm:"column:suppliers_id"`
	AddressID   string `json:"address_id" gorm:"column:address_id"`
}

type ISuppliersRepository interface {
	New(ctx context.Context, sup *Suppliers, addr *Addresses) error
	GetAll(ctx context.Context) ([]Suppliers, error)
}
