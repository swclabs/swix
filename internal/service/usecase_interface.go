package service

import "example/swiftcart/internal/schema"

type IAuth interface {
	SignUp(req schema.SignUpRequest) error
	SignIn(req schema.SignInRequest) (string, error)
}
