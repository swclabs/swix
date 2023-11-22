package service

import (
	"swclabs/swiftcart/internal/schema"
)

// IAccountManagement : Module Account Management with use-case
// Actor: admin & customer (user)
type IAccountManagement interface {
	SignUp(req *schema.SignUpRequest) error
	Login(req *schema.LoginRequest) (string, error)
	UserInfo(email string) (*schema.UserInfo, error)
	ForgetPassword(email string) error
	UpdateUserInfo(req *schema.UserUpdate) error
	UploadAvatar() error
}

// IOrderManagement : Module Order Management
// Actor: admin
type IOrderManagement interface {
}

// IProductManagement : Module Product Management
// Actor: Admin
type IProductManagement interface {
}

// IPurchaseService : Module Purchase
// Actor: Admin & Customer (User)
type IPurchaseService interface {
	AddToCart()
	GetOrders()
	GetCartItems()
	AddCartInfo()
}

// IProductService : Module Product interactions
// Actor: Admin & Customer (User)
type IProductService interface {
	GetProducts()
	SearchProduct()
	GetByCategory()
	SortBy()
	GetProductInfo()
	Like()
	Comment()
}

// IPaymentService : Module Payment
// Actor: Admin & Customer (User)
type IPaymentService interface {
	GetPayments()
}

// IDeliveryService : Module Delivery
// Actor: Admin & Customer (User)
type IDeliveryService interface {
	GetDeliveryInfo()
}
