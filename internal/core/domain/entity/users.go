package entity

// Users :Table users
type Users struct {
	ID          int64  `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Image       string `json:"image" db:"image"`
}

// UserAddress :Table user_address
type UserAddress struct {
	UserID    string `json:"user_id" db:"user_id"`
	AddressID string `json:"address_uuid" db:"address_uuid"`
}