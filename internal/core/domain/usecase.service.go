package domain

import "mime/multipart"

// IAccountManagementService : Module Account Management with use-case
// Actor: admin & customer (user)
type IAccountManagementService interface {
	SignUp(req *SignUpRequest) error
	Login(req *LoginRequest) (string, error)
	CheckLoginEmail(email string) error
	UserInfo(email string) (*UserInfo, error)
	UpdateUserInfo(req *UserUpdate) error
	UploadAvatar(email string, fileHeader *multipart.FileHeader) error
	OAuth2SaveUser(req *OAuth2SaveUser) error
	UpdateAddress(data *Addresses) error
}

// IProductService : Module Product interactions
// Actor: Admin & Customer (User)
type IProductService interface {
	// GetAll()
	// SearchProduct()
	// GetByCategory()
	// SortBy()
	// GetProductInfo()
	// Like()
	// Comment()
}

// IPurchasingService : Module Purchasing
// Actor: Admin & Customer (User)
type IPurchasingService interface {
	// AddToCart()
	// GetOrders()
	// GetCartItems()
	// AddCartInfo()
}

// IProductManagementService : Module Product Management
// Actor: Admin
type IProductManagementService interface {
	InsertCategory(ctg *Categories) error
	UploadImage(Id string, fileHeader *multipart.FileHeader) error
	UploadProduct(img *multipart.FileHeader, products *ProductRequest) error
}

// IOrderManagementService : Module Order Management
// Actor: admin
type IOrderManagementService interface {
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

type ICommonService interface {
	HealthCheck() HealthCheckResponse
	WorkerCheck(num int64) error
}
