package dtos

import "time"

type CreateCoupon struct {
	Status      string `json:"status" validate:"required"`
	MaxUse      int    `json:"max_use" validate:"required"`
	Discount    int    `json:"discount" validate:"required"`
	Description string `json:"description" validate:"required"`
	ExpiredAt   string `json:"expired_at" validate:"required"`
}

type Coupon struct {
	ID          int64     `json:"id"`
	Code        string    `json:"code"`
	Discount    int       `json:"discount"`
	ExpiredAt   time.Time `json:"expired_at"`
	Description string    `json:"description"`
}
