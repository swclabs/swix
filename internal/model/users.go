package model

type User struct {
	UserID      int64  `json:"user_id" gorm:"user_id"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Image       string `json:"image" gorm:"column:image"`
}
