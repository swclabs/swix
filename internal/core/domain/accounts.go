package domain

// Account table
type Account struct {
	Username  string `json:"username" db:"username"`
	Role      string `json:"role" db:"role"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created"`
	Type      string `json:"type" db:"type"`
}
