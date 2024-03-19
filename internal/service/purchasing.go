package service

import "github.com/swclabs/swipe-api/internal/domain"

type Purchasing struct {}

func NewPurchasingService() domain.IPurchasingService {
	return &Purchasing{}
}
