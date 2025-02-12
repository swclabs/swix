// Package entity Categories entities
package entity

// Category Table
type Category struct {
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description" validate:"required"`
}
