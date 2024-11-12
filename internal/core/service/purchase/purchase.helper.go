package purchase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/repos/inventories"
	"github.com/swclabs/swipex/internal/core/repos/orders"
	"github.com/swclabs/swipex/pkg/utils"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

func (p *Purchase) genUUID(ctx context.Context, orderRepo orders.IOrders) string {
	uuid := utils.GenOrderCode(16)
	for {
		_order, err := orderRepo.GetByUUID(ctx, uuid)
		if err != nil || _order.UUID == "" {
			break
		}
		uuid = utils.GenOrderCode(16)
	}
	return uuid
}

func (p *Purchase) calculateTotalAmount(
	ctx context.Context,
	inventory inventories.IInventories,
	order dtos.OrderForm,
) (totalAmount decimal.Decimal, listTotaAmount []decimal.Decimal, err error) {

	totalAmount = decimal.NewFromInt(0)
	listTotaAmount = []decimal.Decimal{}

	for _, product := range order.Product {
		code := strings.Split(product.Code, "#")
		if len(code) != 2 {
			return decimal.NewFromInt(0), nil, fmt.Errorf("invalid product code : %s", product.Code)
		}

		id, err := strconv.ParseInt(code[1], 10, 64)
		if err != nil {
			return decimal.NewFromInt(0), nil, err
		}

		inven, err := inventory.GetByID(ctx, id)
		if err != nil {
			return decimal.NewFromInt(0), nil, err
		}

		totalAmount = totalAmount.Add(inven.Price.Mul(decimal.NewFromInt32(int32(product.Quantity))))
		listTotaAmount = append(listTotaAmount,
			inven.Price.Mul(decimal.NewFromInt32(int32(product.Quantity))))
	}

	return totalAmount, listTotaAmount, nil
}

func (p *Purchase) saveProductOrder(
	ctx context.Context,
	orderRepo orders.IOrders,
	orderID int64,
	order dtos.OrderForm,
	listTotalAmount []decimal.Decimal,
) error {

	for idx, product := range order.Product {

		code := strings.Split(product.Code, "#")
		if len(code) != 2 {
			return fmt.Errorf("invalid product code: %s", product.Code)
		}

		id, err := strconv.ParseInt(code[1], 10, 64)
		if err != nil {
			return err
		}

		if err := orderRepo.InsertProduct(ctx, entity.ProductInOrder{
			OrderID:     orderID,
			InventoryID: id,
			Quantity:    product.Quantity,
			TotalAmount: listTotalAmount[idx],
		}); err != nil {
			return err
		}
	}

	return nil
}

func (p *Purchase) useCoupon(
	ctx context.Context, totalAmount decimal.Decimal,
	order dtos.OrderForm, couponCode string) (newTotalAmount decimal.Decimal, err error) {

	if couponCode == "" {
		return totalAmount, nil
	}

	coupon, err := p.Coupon.GetByCode(ctx, order.CouponCode)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return decimal.NewFromInt32(0), err
	}

	if coupon != nil {
		discount := decimal.NewFromInt32(int32(coupon.Discount)).Div(decimal.NewFromInt32(100)).Mul(totalAmount)
		newAmount := totalAmount.Copy().Sub(discount)
		return newAmount, nil
	}
	return totalAmount, nil
}

func (p *Purchase) getListOrder(ctx context.Context, orders []entity.Order) ([]dtos.OrderInfo, error) {
	var orderInfo []dtos.OrderInfo

	for _, order := range orders {

		// GetByUserID products by order code
		order, err := p.GetOrderByCode(ctx, order.UUID)
		if err != nil {
			return nil, fmt.Errorf("error getting order by UUID: %w", err)
		}

		// Merge product and order schema
		orderInfo = append(orderInfo, *order)
	}

	return orderInfo, nil
}
