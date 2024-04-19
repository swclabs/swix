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
	GetProductsLimit(ctx context.Context, limit int) ([]Products, error)
	InsertCategory(ctx context.Context, ctg *Categories) error
	UploadProductImage(ctx context.Context, Id int, fileHeader *multipart.FileHeader) error
	UploadProduct(ctx context.Context, fileHeader *multipart.FileHeader, products ProductRequest) error
	GetSuppliersLimit(ctx context.Context, limit int) ([]Suppliers, error)
	InsertSuppliers(ctx context.Context, supplierReq SuppliersRequest) error
}

// IPurchasingService : Module Purchasing
// Actor: Admin & Customer (User)
type IPurchasingService interface {
	AddToCart(ctx context.Context, cart CartInfo)
	GetCart(ctx context.Context, limit int) ([]Carts, error)
	GetOrders(ctx context.Context, limit int) ([]Orders, error)
	InsertOrders(ctx context.Context, order Orders) error
}

type IPostsService interface {
	GetNewsletter(ctx context.Context, limit int) ([]Newsletters, error)
	GetHomeBanner(ctx context.Context, limit int) ([]HomeBanners, error)
	UploadNewsletter(ctx context.Context, news Newsletter, fileHeader *multipart.FileHeader) error
	UploadHomeBanner(ctx context.Context, data HomeBanners, fileHeader *multipart.FileHeader) error
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
