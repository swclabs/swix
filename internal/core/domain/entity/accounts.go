// Package entity Account entities
package entity

import "time"

// Account table
type Account struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Role      string    `json:"role" db:"role"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
