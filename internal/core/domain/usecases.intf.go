package domain

import (
	"context"
	"mime/multipart"
)

// IAccountManagementService : Module for Account Management with use-cases.
// Actor: Admin & Customer (User)
type IAccountManagementService interface {
	// SignUp registers a new user.
	// ctx is the context to manage the request's lifecycle.
	// req contains the sign-up request details.
	// Returns an error if any issues occur during the sign-up process.
	SignUp(ctx context.Context, req *SignUpReq) error

	// Login authenticates a user and returns a token.
	// ctx is the context to manage the request's lifecycle.
	// req contains the login request details.
	// Returns a token string and an error if any issues occur during the login process.
	Login(ctx context.Context, req *LoginReq) (string, error)

	// CheckLoginEmail checks if the email is already registered.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to check.
	// Returns an error if any issues occur during the check process.
	CheckLoginEmail(ctx context.Context, email string) error

	// UserInfo retrieves user information based on email.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to retrieve user information for.
	// Returns a pointer to the UserInfo object and an error if any issues occur during the retrieval process.
	UserInfo(ctx context.Context, email string) (*UserInfo, error)

	// UpdateUserInfo updates the user information.
	// ctx is the context to manage the request's lifecycle.
	// req contains the updated user information details.
	// Returns an error if any issues occur during the update process.
	UpdateUserInfo(ctx context.Context, req *UserUpdate) error

	// UploadAvatar uploads a user's avatar.
	// email is the email address of the user.
	// fileHeader contains the file header of the avatar to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadAvatar(email string, fileHeader *multipart.FileHeader) error

	// OAuth2SaveUser saves user information from an OAuth2 login.
	// ctx is the context to manage the request's lifecycle.
	// req contains the OAuth2 user information details.
	// Returns an error if any issues occur during the save process.
	OAuth2SaveUser(ctx context.Context, req *OAuth2SaveUser) error

	// UploadAddress uploads a user's address.
	// ctx is the context to manage the request's lifecycle.
	// data contains the address details to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadAddress(ctx context.Context, data *Addresses) error
}

// IProductService : Module for Product interactions.
// Actor: Admin & Customer (User)
type IProductService interface {
	// GetCategoriesLimit retrieves a list of categories with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of categories to retrieve.
	// Returns a slice of Categories objects and an error if any issues occur during the retrieval process.
	GetCategoriesLimit(ctx context.Context, limit string) ([]Categories, error)

	// GetProductsLimit retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductRes objects and an error if any issues occur during the retrieval process.
	GetProductsLimit(ctx context.Context, limit int) ([]ProductRes, error)

	// InsertCategory adds a new category to the database.
	// ctx is the context to manage the request's lifecycle.
	// ctg is a pointer to the Categories object to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertCategory(ctx context.Context, ctg *Categories) error

	// UploadProductImage uploads images for a product.
	// ctx is the context to manage the request's lifecycle.
	// Id is the ID of the product.
	// fileHeader contains the file headers of the images to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadProductImage(ctx context.Context, Id int, fileHeader []*multipart.FileHeader) error

	// UploadProduct adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// products contains the product details to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	UploadProduct(ctx context.Context, products ProductReq) (int64, error)

	// GetSuppliersLimit retrieves a list of suppliers with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of suppliers to retrieve.
	// Returns a slice of Suppliers objects and an error if any issues occur during the retrieval process.
	GetSuppliersLimit(ctx context.Context, limit int) ([]Suppliers, error)

	// InsertSuppliers adds a new supplier to the database.
	// ctx is the context to manage the request's lifecycle.
	// supplierReq contains the supplier details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertSuppliers(ctx context.Context, supplierReq SuppliersReq) error

	// InsertIntoWarehouse adds a product to the warehouse inventory.
	// ctx is the context to manage the request's lifecycle.
	// product contains the warehouse product details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertIntoWarehouse(ctx context.Context, product WarehouseReq) error

	// GetProductsInWarehouse retrieves product details from the warehouse.
	// ctx is the context to manage the request's lifecycle.
	// productID, ram, ssd, and color specify the product attributes to retrieve.
	// Returns a pointer to the WarehouseRes object and an error if any issues occur during the retrieval process.
	GetProductsInWarehouse(ctx context.Context, productID, ram, ssd, color string) (*WarehouseRes, error)
}

// IPurchaseService : Module for Purchasing.
// Actor: Admin & Customer (User)
type IPurchaseService interface {
	// AddToCart adds a product to the shopping cart.
	// ctx is the context to manage the request's lifecycle.
	// cart contains the cart information to be added.
	AddToCart(ctx context.Context, cart CartInfo)

	// GetCart retrieves the shopping cart with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of cart items to retrieve.
	// Returns a slice of Carts objects and an error if any issues occur during the retrieval process.
	GetCart(ctx context.Context, limit int) ([]Carts, error)

	// GetOrders retrieves orders with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of orders to retrieve.
	// Returns a slice of Orders objects and an error if any issues occur during the retrieval process.
	GetOrders(ctx context.Context, limit int) ([]Orders, error)

	// InsertOrders adds a new order to the database.
	// ctx is the context to manage the request's lifecycle.
	// order contains the order details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertOrders(ctx context.Context, order Orders) error
}

// IPostsService : Module for managing posts.
// Actor: Admin & Customer
type IPostsService interface {
	// UploadCollections uploads a new collection.
	// ctx is the context to manage the request's lifecycle.
	// banner contains the collection details to be uploaded.
	// Returns id of collection was uploaded and error if any issues occur during the upload process.
	UploadCollections(ctx context.Context, banner CollectionType) (int64, error)

	// UploadCollectionsImage uploads a new image of collection.
	// ctx is the context to manage the request's lifecycle.
	// cardBannerId contains the id of collection to be uploaded.
	// fileHeader is  the header of the file to be uploaded
	// Returns an error if any issues occur during the upload process.
	UploadCollectionsImage(ctx context.Context, cardBannerId string, fileHeader *multipart.FileHeader) error

	// SlicesOfCollections return a slices of collection.
	// ctx is the context to manage the request's lifecycle.
	// cardBannerId contains the id of collection to be returns.
	// limit is the maximum number of Collection to retrieve.
	// Returns an error if any issues occur during the upload process.
	SlicesOfCollections(ctx context.Context, position string, limit int) (*Collections, error)
}

// IOrderManagementService : Module for Order Management.
// Actor: Admin
type IOrderManagementService interface {
	// Define methods for order management as needed.
}

// IPaymentService : Module for Payment.
// Actor: Admin & Customer (User)
type IPaymentService interface {
	// GetPayments retrieves payment information.
	// ctx is the context to manage the request's lifecycle.
	GetPayments(ctx context.Context)
}

// IDeliveryService : Module for Delivery.
// Actor: Admin & Customer (User)
type IDeliveryService interface {
	// GetDeliveryInfo retrieves delivery information.
	// ctx is the context to manage the request's lifecycle.
	GetDeliveryInfo(ctx context.Context)
}

// ICommonService : Common utility methods for the service.
// Actor: System
type ICommonService interface {
	// HealthCheck performs a health check on the service.
	// ctx is the context to manage the request's lifecycle.
	// Returns a HealthCheckResponse object with the health check status.
	HealthCheck(ctx context.Context) HealthCheckResponse

	// WorkerCheck checks the status of a worker.
	// ctx is the context to manage the request's lifecycle.
	// num specifies the worker number to check.
	// Returns an error if any issues occur during the check process.
	WorkerCheck(ctx context.Context, num int64) error
}
