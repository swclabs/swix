package service

import (
	"context"
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
)

type ProductService struct {
	newsletter domain.INewsletterRepository
	categories domain.ICategoriesRepository
	products   domain.IProductRepository
}

func NewProductService() domain.IProductService {
	return &ProductService{
		newsletter: repo.NewNewsletter(),
		categories: repo.NewCategories(),
		products:   repo.NewProducts(),
	}
}

func (s *ProductService) GetNewsletter(ctx context.Context, limit int) ([]domain.Newsletters, error) {
	return s.newsletter.Get(ctx, limit)
}

func (s *ProductService) GetHomeBanner(ctx context.Context) ([]domain.HomeBanners, error) {
	// TODO:
	return nil, nil
}

func (s *ProductService) GetAccessory(ctx context.Context) ([]domain.Accessory, error) {
	// TODO:
	return nil, nil
}

func (s *ProductService) GetCategoriesLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	return s.categories.GetLimit(ctx, limit)
}

func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]domain.Products, error) {
	return s.products.GetLitmit(ctx, limit)
}
