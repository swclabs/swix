// Package domain Account entities
package domain

import "time"

// Account table
type Account struct {
	Username  string    `json:"username" db:"username"`
	Role      string    `json:"role" db:"role"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Type      string    `json:"type" db:"type"`
}
