package model

type Users struct {
	Username          string `json:"username"`
	HashedPassword    string `json:"hashed_password" gorm:"column:hashed_password"`
	FullName          string `json:"full_name" gorm:"column:full_name"`
	Email             string `json:"email" gorm:"column:email"`
	PasswordChangedAt string `json:"password_changed_at" gorm:"column:password_changed_at"`
	CreatedAt         string `json:"created_at" gorm:"column:created_at"`
}

type Account struct {
	ID        string `json:"id" gorm:"column:id"`
	Owner     string `json:"owner" gorm:"column:owner"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
}
