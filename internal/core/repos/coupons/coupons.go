package coupons

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(New)

func New(conn db.IDatabase) ICoupons {
	return &Coupon{conn}
}

type Coupon struct {
	db db.IDatabase
}

// GetByCode implements ICoupons.
func (c *Coupon) GetByCode(ctx context.Context, code string) (*entity.Coupons, error) {
	rows, err := c.db.Query(ctx, getByCode, code)
	if err != nil {
		return nil, err
	}
	coupon, err := db.CollectRow[entity.Coupons](rows)
	if err != nil {
		return nil, err
	}
	return &coupon, nil
}

// GetByUser implements ICoupons.
func (c *Coupon) GetByUser(ctx context.Context, userID int64) ([]entity.CouponsUsed, error) {
	rows, err := c.db.Query(ctx, getByUser, userID)
	if err != nil {
		return nil, err
	}
	coupons, err := db.CollectRows[entity.CouponsUsed](rows)
	if err != nil {
		return nil, err
	}
	return coupons, nil
}

// GetAll implements ICoupons.
func (c *Coupon) GetAll(ctx context.Context) ([]entity.Coupons, error) {
	rows, err := c.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	coupons, err := db.CollectRows[entity.Coupons](rows)
	if err != nil {
		return nil, err
	}
	return coupons, nil
}

// Create implements ICoupons.
func (c *Coupon) Create(ctx context.Context, coupon entity.Coupons) error {
	return c.db.SafeWrite(ctx, insert, coupon.Code, coupon.Discount,
		coupon.Status, coupon.Used, coupon.MaxUse, coupon.Description, coupon.ExpiredAt)
}

// Use implements ICoupons.
func (c *Coupon) Use(ctx context.Context, couponInfo entity.CouponsUsed) error {
	return c.db.SafeWrite(ctx, useCoupons, couponInfo.CouponCode, couponInfo.UserID, couponInfo.CouponCode,
		couponInfo.OrderID, couponInfo.UsedAt)
}
