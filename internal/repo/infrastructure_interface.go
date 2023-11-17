package repo

import (
	"example/swiftcart/internal/model"
	"example/swiftcart/internal/schema"
)

type IUsers interface {
	GetByEmail(email string) (*model.User, error)
	Insert(usr *model.User) error
	Infor(email string) (*schema.InforResponse, error)
}

type IAccounts interface {
	GetByEmail(email string) (*model.Account, error)
	Insert(acc *model.Account) error
}
