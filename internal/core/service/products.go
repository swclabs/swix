package service

import (
	"context"
	"swclabs/swipe-api/internal/core/domain"
)

type ProductService struct{}

func NewProductService() domain.IProductService {
	return &ProductService{}
}

func (s *ProductService) GetNewsletter(ctx context.Context) ([]domain.Newsletter, error) {
	// TODO:
	return nil, nil
}

func (s *ProductService) GetHomeBanner(ctx context.Context) ([]domain.HomeBanners, error) {
	// TODO:
	return nil, nil
}

func (s *ProductService) GetAccessory(ctx context.Context) ([]domain.Accessory, error) {
	// TODO:
	return nil, nil
}
