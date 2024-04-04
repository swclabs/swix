package domain

import (
	"context"
	"mime/multipart"
)

// IAccountManagementService : Module Account Management with use-case
// Actor: admin & customer (user)
type IAccountManagementService interface {
	SignUp(ctx context.Context, req *SignUpRequest) error
	Login(ctx context.Context, req *LoginRequest) (string, error)
	CheckLoginEmail(ctx context.Context, email string) error
	UserInfo(ctx context.Context, email string) (*UserInfo, error)
	UpdateUserInfo(ctx context.Context, req *UserUpdate) error
	UploadAvatar(email string, fileHeader *multipart.FileHeader) error
	OAuth2SaveUser(ctx context.Context, req *OAuth2SaveUser) error
	UploadAddress(ctx context.Context, data *Addresses) error
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

	GetAccessory(ctx context.Context) ([]Accessory, error)
	GetCategoriesLimit(ctx context.Context, limit string) ([]Categories, error)
	GetNewsletter(ctx context.Context, limit int) ([]Newsletter, error)
	GetHomeBanner(ctx context.Context) ([]HomeBanners, error)
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
	InsertCategory(ctx context.Context, ctg *Categories) error
	UploadImage(Id string, fileHeader *multipart.FileHeader) error
	UploadProduct(img *multipart.FileHeader, products *ProductRequest) error
	UploadNewsletter(ctx context.Context, news Newsletter, fileHeader *multipart.FileHeader) error
	UploadHomeBanner(ctx context.Context, data *HomeBanners) error
	GetAllSuppliers(ctx context.Context) ([]Suppliers, error)
}

// IOrderManagementService : Module Order Management
// Actor: admin
type IOrderManagementService interface {
}

// IPaymentService : Module Payment
// Actor: Admin & Customer (User)
type IPaymentService interface {
	GetPayments(ctx context.Context)
}

// IDeliveryService : Module Delivery
// Actor: Admin & Customer (User)
type IDeliveryService interface {
	GetDeliveryInfo(ctx context.Context)
}

type ICommonService interface {
	HealthCheck(ctx context.Context) HealthCheckResponse
	WorkerCheck(ctx context.Context, num int64) error
}
