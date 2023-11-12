package repository

import "example/komposervice/internal/model"

type IUsers interface {
	GetByEmail(email string) (*model.Users, error)
	Create(_usr *model.Users) error
	Empty() bool
}
