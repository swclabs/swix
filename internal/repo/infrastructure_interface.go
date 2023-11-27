package repo

import (
	"swclabs/swiftcart/internal/model"
	"swclabs/swiftcart/internal/schema"
)

type IUsers interface {
	GetByEmail(email string) (*model.User, error)
	Insert(usr *model.User) error
	Info(email string) (*schema.UserInfo, error)
	SaveInfo(user *model.User) error
	OAuth2SaveInfo(user *model.User) error
}

type IAccounts interface {
	GetByEmail(email string) (*model.Account, error)
	Insert(acc *model.Account) error
	SaveInfo(acc *model.Account) error
}

type ICarts interface {
	Add(productID int64) error
	AddMany(products []int64) error
	GetCartByUserID(userId int64) (*schema.CartInfo, error)
	RemoveProduct(productID int64) error
}

type IProductInCart interface {
	GetByCartID(cartID int64) ([]model.ProductInCart, error)
	AddProduct(product *model.ProductInCart) error
	RemoveProduct(productID, cartID int64) error
	Save(product *model.ProductInCart) error
}
