// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"fmt"
	"log"
	"strings"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/repository/carts"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/repository/orders"
	"swclabs/swipecore/internal/core/repository/users"
	"swclabs/swipecore/pkg/infra/db"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Purchase struct for purchase service
type Purchase struct {
	Order orders.IOrdersRepository
	Cart  carts.ICartRepository
	User  users.IUserRepository
}

// New creates a new Purchase object
func New(
	order orders.IOrdersRepository,
	cart carts.ICartRepository,
	user users.IUserRepository,
) IPurchaseService {
	return &Purchase{
		Cart:  cart,
		Order: order,
		User:  user,
	}
}

// GetOrdersByUserID implements IPurchaseService.
func (p *Purchase) GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderSchema, error) {
	// Get orders by user ID
	orders, err := p.Order.Get(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	// Get user by user ID
	user, err := p.User.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var orderSchema []dtos.OrderSchema

	for _, order := range orders {
		// Get products by order ID
		products, err := p.Order.GetProductByOrderID(ctx, order.ID)
		if err != nil {
			return nil, err
		}
		var productSchema []dtos.ProductOrderSchema
		for _, product := range products {
			productSchema = append(productSchema, dtos.ProductOrderSchema{
				ID:           product.ID,
				OrderID:      product.OrderID,
				CurrencyCode: product.CurrencyCode,
				InventoryID:  product.InventoryID,
				Quantity:     product.Quantity,
				TotalAmount:  product.TotalAmount.String(),
			})
		}
		// Merge product and order schema
		orderSchema = append(orderSchema, dtos.OrderSchema{
			ID:        order.ID,
			UUID:      order.UUID,
			Status:    order.Status,
			Products:  productSchema,
			UserEmail: user.Email,
			UserID:    user.ID,
			Username:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		})
	}
	return orderSchema, nil
}

// DeleteItemFromCart implements IPurchaseService.
func (p *Purchase) DeleteItemFromCart(ctx context.Context, userID int64, inventoryID int64) error {
	return p.Cart.RemoveItem(ctx, userID, inventoryID)
}

// AddToCart implements IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart dtos.CartInsert) error {
	return p.Cart.Insert(ctx, cart.UserID, cart.InventoryID, cart.Quantity)
}

// GetCart implements IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userID int64, limit int) (*dtos.CartSlices, error) {
	cartItems, err := p.Cart.GetCartByUserID(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	var cartSchema dtos.CartSlices
	cartSchema.UserID = userID
	for _, item := range cartItems {

		cartSchema.Products = append(cartSchema.Products, dtos.CartSchema{
			Quantity: item.Quantity,
		})
	}
	return &cartSchema, nil
}

// CreateOrders implements IPurchaseService.
func (p *Purchase) CreateOrders(ctx context.Context, createOrder dtos.CreateOrderSchema) (string, error) {
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return "", err
	}
	var (
		_uuid              = strings.ReplaceAll(uuid.New().String(), "_", " ")
		inventoryRepo      = inventories.New(tx)
		orderRepo          = orders.New(tx)
		totalAmount        decimal.Decimal
		productTotalAmount []decimal.Decimal
	)

	for _, product := range createOrder.Products {
		inven, err := inventoryRepo.GetByID(ctx, product.InventoryID)
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
		totalAmount = totalAmount.Add(inven.Price)
		productTotalAmount = append(
			productTotalAmount,
			inven.Price.Mul(decimal.NewFromInt32(int32(product.Quantity))))
	}
	orderID, err := orderRepo.Create(ctx, entity.Orders{
		UUID:        _uuid,
		UserID:      createOrder.UserID,
		Status:      "pending",
		TotalAmount: totalAmount,
	})
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return "", err
	}
	for idx, product := range createOrder.Products {
		if err := orderRepo.InsertProduct(ctx, entity.ProductInOrder{
			OrderID:     orderID,
			InventoryID: product.InventoryID,
			Quantity:    product.Quantity,
			TotalAmount: productTotalAmount[idx],
		}); err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
	}
	return _uuid, tx.Commit(ctx)
}
