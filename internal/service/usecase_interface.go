package service

import "example/swiftcart/internal/schema"

type IAccountManagement interface {
	SignUp(req *schema.SignUpRequest) error
	Login(req *schema.LoginRequest) (string, error)
	Info(email string) (*schema.UserInfo, error)
}
