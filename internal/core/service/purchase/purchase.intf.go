// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"swclabs/swix/internal/core/domain/dtos"
)

// IPurchaseService : Module for Purchasing.
// Actor: Admin & Customer (Users)
type IPurchaseService interface {
	// AddToCart adds a product to the shopping cart.
	// ctx is the context to manage the request's lifecycle.
	// cart contains the cart information to be added.
	AddToCart(ctx context.Context, cart dtos.CartInsert) error

	// GetCart retrieves the shopping cart with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of cart items to retrieve.
	// userId is the user ID of cart item to retrieve.
	// Returns a slice of Carts objects and an error if any issues occur during the retrieval process.
	GetCart(ctx context.Context, userID int64, limit int) (*dtos.CartSlices, error)

	// CreateOrders creates a new order.
	// ctx is the context to manage the request's lifecycle.
	// createOrder contains the order information to be created.
	// Returns the UUID of the newly created order and an error if any issues occur during the creation process.
	CreateOrders(ctx context.Context, createOrder dtos.CreateOrderSchema) (string, error)

	// DeleteItemFromCart deletes an item from the shopping cart.
	// ctx is the context to manage the request's lifecycle.
	// userID is the user ID of the cart item to delete.
	// inventoryID is the inventory ID of the cart item to delete.
	// Returns an error if any issues occur during the deletion process.
	DeleteItemFromCart(ctx context.Context, cartID int64) error

	// GetOrdersByUserID retrieves orders by user ID with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// userID is the user ID of the orders to retrieve.
	// limit is the maximum number of orders to retrieve.
	// Returns a slice of OrderSchema objects and an error if any issues occur during the retrieval process.
	GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderSchema, error)
}
