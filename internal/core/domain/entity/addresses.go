// Package entity Addresses entities
package entity

// Address Table
type Address struct {
	ID       int64  `json:"id" db:"id"`
	UserID   int64  `json:"user_id" db:"user_id"`
	City     string `json:"city" db:"city"`
	Ward     string `json:"ward" db:"ward"`
	District string `json:"district" db:"district"`
	Street   string `json:"street" db:"street"`
}
