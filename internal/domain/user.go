package domain

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

// User Repository interface
// implement at /internal/repo/user.go
type IUserRepository interface {
	GetByEmail(email string) (*User, error)
	Insert(usr *User) error
	Info(email string) (*UserInfo, error)
	SaveInfo(user *User) error
	OAuth2SaveInfo(user *User) error
	SignUp(user *User, password string) error
	SaveOAuth2(data *User) error
}

// SignUpRequest schema
type SignUpRequest struct {
	Email       string `json:"email" validate:"email,required"`
	PhoneNumber string `json:"phone_number" validate:"number,required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// SignUpResponse schema
type SignUpResponse struct {
	Success bool   `json:"success" validate:"required"`
	Msg     string `json:"msg" validate:"required"`
}

// LoginRequest schema
type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse schema
type LoginResponse struct {
	Success bool   `json:"success" validate:"required"`
	Token   string `json:"token" validate:"required"`
	Email   string `json:"email" validate:"email,required"`
}

// UserInfo schema
type UserInfo struct {
	Id          int64  `json:"id" validate:"required"`
	Email       string `json:"email" validate:"email,required"`
	PhoneNumber string `json:"phone_number" validate:"number,required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Role        string `json:"role" validate:"required"`
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
