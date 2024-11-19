// Package purchase implements the purchase interface
package purchase

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/x/ghn"
)

// IPurchase : Module for Purchasing.
// Actor: Admin & Customer (Users)
type IPurchase interface {
	// AddToCart adds a product to the shopping cart.
	// ctx is the context to manage the request's lifecycle.
	// cart contains the cart information to be added.
	AddToCart(ctx context.Context, cart dtos.CartInsertDTO) error

	// GetCart retrieves the shopping cart with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of cart items to retrieve.
	// userId is the user ID of cart item to retrieve.
	// Returns a slice of Carts objects and an error if any issues occur during the retrieval process.
	GetCart(ctx context.Context, userID int64, limit int) (*dtos.Carts, error)

	// CreateOrders creates a new order.
	// ctx is the context to manage the request's lifecycle.
	// createOrder contains the order information to be created.
	// Returns the UUID of the newly created order and an error if any issues occur during the creation process.
	CreateOrders(ctx context.Context, userID int64, createOrder dtos.Order) (string, error)

	CreateOrderForm(ctx context.Context, order dtos.OrderForm) (string, error)

	// DeleteItemFromCart deletes an item from the shopping cart.
	// ctx is the context to manage the request's lifecycle.
	// userID is the user ID of the cart item to delete.
	// inventoryID is the inventory ID of the cart item to delete.
	// Returns an error if any issues occur during the deletion process.
	DeleteItemFromCart(ctx context.Context, inventoryID int64, userID int64) error

	// GetOrdersByUserID retrieves orders by user ID with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// userID is the user ID of the orders to retrieve.
	// limit is the maximum number of orders to retrieve.
	// Returns a slice of OrderSchema objects and an error if any issues occur during the retrieval process.
	GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderInfo, error)

	GetUsersByAdmin(ctx context.Context, limit int) ([]dtos.OrderInfo, error)

	GetOrderByCode(ctx context.Context, orderCode string) (*dtos.OrderInfo, error)

	UpdateOrderStatus(ctx context.Context, orderCode string, status string) error

	DeliveryOrderInfo(ctx context.Context, orderCode string) (*ghn.OrderInfoDTO, error)

	CreateDeliveryOrder(ctx context.Context, shopID int, order ghn.CreateOrderDTO) (*ghn.OrderDTO, error)

	// CreateDeliveryAddress creates a new delivery address.
	// ctx is the context to manage the request's lifecycle.
	// addr contains the delivery address information to be created.
	// Returns an error if any issues occur during the creation process.
	CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error

	// GetDeliveryAddress retrieves delivery addresses by user ID.
	// ctx is the context to manage the request's lifecycle.
	// userID is the user ID of the delivery addresses to retrieve.
	// Returns a slice of Address objects and an error if any issues occur during the retrieval process.
	GetDeliveryAddress(ctx context.Context, userID int64) ([]dtos.Address, error)

	// CreateDelivery creates a new delivery.
	// ctx is the context to manage the request's lifecycle.
	// delivery contains the delivery information to be created.
	// Returns an error if any issues occur during the creation process.
	CreateDelivery(ctx context.Context, delivery dtos.DeliveryBody) error

	// GetDelivery retrieves deliveries by user ID.
	// ctx is the context to manage the request's lifecycle.
	// userID is the user ID of the deliveries to retrieve.
	// Returns a slice of Delivery objects and an error if any issues occur during the retrieval process.
	GetDelivery(ctx context.Context, userID int64) ([]dtos.Delivery, error)

	AddressProvince(ctx context.Context) ([]entity.Province, error)

	AddressWard(ctx context.Context, districtID string) ([]entity.Commune, error)

	AddressDistrict(ctx context.Context, provinceID string) ([]entity.District, error)

	CreateCoupon(ctx context.Context, coupon dtos.CreateCoupon) (code string, err error)

	GetCoupon(ctx context.Context) (coupons []dtos.Coupon, err error)

	DeleteCoupon(ctx context.Context, code string) error
}
