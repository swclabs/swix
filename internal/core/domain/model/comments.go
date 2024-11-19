package model

import "time"

// comment_id as id, user_id, email, first_name, last_name, rating, content, specs, color, created
type Comment struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Email       string    `json:"email" db:"email"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Rating      float64   `json:"rating" db:"rating"`
	Content     string    `json:"content" db:"content"`
	ProductName string    `json:"product_name" db:"product_name"`
	Specs       string    `json:"specs" db:"specs"`
	Color       string    `json:"color" db:"color"`
	Created     time.Time `json:"created" db:"created"`
}
