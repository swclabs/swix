// Package entity Categories entities
package entity

// Categories Table
type Categories struct {
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description" validate:"required"`
}
