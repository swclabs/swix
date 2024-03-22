package service

import (
	"github.com/swclabs/swipe-api/internal/core/domain"
)

type ProductService struct{}

func NewProductService() domain.IProductService {
	return &ProductService{}
}
