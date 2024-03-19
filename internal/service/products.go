package service

import (
	"github.com/swclabs/swipe-api/internal/domain"
)

type ProductService struct{}

func NewProductService() domain.IProductService {
	return &ProductService{}
}
