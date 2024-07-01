package domain

// Users :Table users
type Users struct {
	Id          int64  `json:"id" db:"id"`
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

/*****************************************************************************/

// SignUpSchema schema
type SignUpSchema struct {
	Email       string `json:"email" validate:"email,required"`
	PhoneNumber string `json:"phone_number" validate:"number,required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// SignUpMessage schema
type SignUpMessage struct {
	Success bool   `json:"success" validate:"required"`
	Msg     string `json:"msg" validate:"required"`
}

// LoginSchema schema
type LoginSchema struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

// LoginMessage schema
type LoginMessage struct {
	Success bool   `json:"success" validate:"required"`
	Token   string `json:"token" validate:"required"`
	Email   string `json:"email" validate:"email,required"`
}

// UserSchema schema
type UserSchema struct {
	Id          int64  `json:"id" validate:"required" db:"id"`
	Email       string `json:"email" validate:"email,required" db:"email"`
	PhoneNumber string `json:"phone_number" validate:"number,required" db:"phone_number"`
	FirstName   string `json:"first_name" validate:"required" db:"first_name"`
	LastName    string `json:"last_name" validate:"required" db:"last_name"`
	Image       string `json:"image" validate:"required" db:"image"`
	Username    string `json:"username" validate:"required" db:"username"`
	Role        string `json:"role" validate:"required" db:"role"`
}

// UserUpdate schema
type UserUpdate struct {
	Id          int64  `json:"id" validate:"required"`
	Email       string `json:"email" validate:"email,required"`
	PhoneNumber string `json:"phone_number" validate:"number,required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Username    string `json:"username" validate:"required"`
}

// OAuth2SaveUser schema
type OAuth2SaveUser struct {
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"number,required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Image       string `json:"image" validate:"required"`
}
