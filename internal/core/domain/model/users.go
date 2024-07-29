// Package model contains the domain models for the application.
package model

// Users model is a representation of the users joined table.
type Users struct {
	ID          int64  `json:"id" validate:"required" db:"id"`
	Email       string `json:"email" validate:"email,required" db:"email"`
	PhoneNumber string `json:"phone_number" validate:"number,required" db:"phone_number"`
	FirstName   string `json:"first_name" validate:"required" db:"first_name"`
	LastName    string `json:"last_name" validate:"required" db:"last_name"`
	Image       string `json:"image" validate:"required" db:"image"`
	Username    string `json:"username" validate:"required" db:"username"`
	Role        string `json:"role" validate:"required" db:"role"`
}
