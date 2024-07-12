// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"log"
	"strings"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/carts"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/repository/orders"
	"swclabs/swipecore/pkg/infra/db"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Purchase struct for purchase service
type Purchase struct {
	Order orders.IOrdersRepository
	Cart  carts.ICartRepository
}

// New creates a new Purchase object
func New(
	order orders.IOrdersRepository,
	cart carts.ICartRepository,
) IPurchaseService {
	return &Purchase{
		Cart:  cart,
		Order: order,
	}
}

// DeleteItemFromCart implements domain.IPurchaseService.
func (p *Purchase) DeleteItemFromCart(ctx context.Context, userID int64, inventoryID int64) error {
	return p.Cart.RemoveItem(ctx, userID, inventoryID)
}

// AddToCart implements domain.IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart domain.CartInsert) error {
	return p.Cart.Insert(ctx, cart.UserID, cart.InventoryID, cart.Quantity)
}

// GetCart implements domain.IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userID int64, limit int) (*domain.CartSlices, error) {
	return p.Cart.GetCartByUserID(ctx, userID, limit)
}

// GetOrders implements domain.IPurchaseService.
func (p *Purchase) GetOrders(_ context.Context, _ int) ([]domain.Orders, error) {
	panic("unimplemented")
}

// InsertOrders implements domain.IPurchaseService.
func (p *Purchase) InsertOrders(ctx context.Context, createOrder domain.CreateOrderSchema) (string, error) {
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
	orderID, err := orderRepo.Create(ctx, domain.Orders{
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
		if err := orderRepo.InsertProduct(ctx, domain.ProductInOrder{
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
