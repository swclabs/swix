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

type Purchase struct {
	Order orders.IOrdersRepository
	Cart  carts.ICartRepository
}

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
func (p *Purchase) DeleteItemFromCart(ctx context.Context, userId int64, inventoryId int64) error {
	return p.Cart.RemoveItem(ctx, userId, inventoryId)
}

// AddToCart implements domain.IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart domain.CartInsert) error {
	return p.Cart.Insert(ctx, cart.UserId, cart.InventoryId, cart.Quantity)
}

// GetCart implements domain.IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userId int64, limit int) (*domain.CartSlices, error) {
	return p.Cart.GetCartByUserID(ctx, userId, limit)
}

// GetOrders implements domain.IPurchaseService.
func (p *Purchase) GetOrders(ctx context.Context, limit int) ([]domain.Orders, error) {
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
		inven, err := inventoryRepo.GetById(ctx, product.InventoryId)
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
	orderId, err := orderRepo.Create(ctx, domain.Orders{
		Uuid:        _uuid,
		UserId:      createOrder.UserId,
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
			OrderId:     orderId,
			InventoryId: product.InventoryId,
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
