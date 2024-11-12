package coupons

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
)

type ICoupons interface {
	Create(ctx context.Context, coupon entity.Coupons) error
	GetAll(ctx context.Context) ([]entity.Coupons, error)
	GetByCode(ctx context.Context, code string) (*entity.Coupons, error)
	GetByUser(ctx context.Context, userID int64) ([]entity.CouponsUsed, error)
	Use(ctx context.Context, couponInfo entity.CouponsUsed) error
}
