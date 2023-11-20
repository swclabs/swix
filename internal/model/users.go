package model

// User : Table users
type User struct {
	UserID      int64  `json:"id" gorm:"column:id"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Image       string `json:"image" gorm:"column:image"`
}

// UserAddress :Table user_address
type UserAddress struct {
	UserID    string `json:"user_id" gorm:"column:user_id"`
	AddressID string `json:"address_id" gorm:"column:address_id"`
}
