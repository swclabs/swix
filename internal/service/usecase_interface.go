package service

import "example/swiftcart/internal/schema"

type IUser interface {
	SignUp(req *schema.SignUpRequest) error
	Login(req *schema.LoginRequest) (string, error)
	Infor(email string) (*schema.UserInfor, error)
}
