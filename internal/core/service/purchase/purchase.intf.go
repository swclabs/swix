package purchase

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// IPurchaseService : Module for Purchasing.
// Actor: Admin & Customer (User)
type IPurchaseService interface {
	// AddToCart adds a product to the shopping cart.
	// ctx is the context to manage the request's lifecycle.
	// cart contains the cart information to be added.
	AddToCart(ctx context.Context, cart domain.CartInsertReq) error

	// GetCart retrieves the shopping cart with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of cart items to retrieve.
	// userId is the user ID of cart item to retrieve.
	// Returns a slice of Carts objects and an error if any issues occur during the retrieval process.
	GetCart(ctx context.Context, userId int64, limit int) (*domain.CartSchema, error)

	// GetOrders retrieves orders with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of orders to retrieve.
	// Returns a slice of Orders objects and an error if any issues occur during the retrieval process.
	GetOrders(ctx context.Context, limit int) ([]domain.Orders, error)

	// InsertOrders adds a new order to the database.
	// ctx is the context to manage the request's lifecycle.
	// order contains the order details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertOrders(ctx context.Context, order domain.Orders) error

	DeleteItemFromCart(ctx context.Context, userId int64, inventoryId int64) error
}
