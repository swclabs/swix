package service

import (
	"context"
	"mime/multipart"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
)

type ProductManagement struct {
	Category domain.ICategoriesRepository
}

func NewProductManagement() domain.IProductManagementService {
	return &ProductManagement{
		Category: repo.NewCategories(),
	}
}

func (product *ProductManagement) InsertCategory(ctx context.Context, ctg *domain.Categories) error {
	// call repository layer
	return product.Category.New(ctx, ctg)
}

func (product *ProductManagement) UploadImage(Id string, fileHeader *multipart.FileHeader) error {
	return nil
}

func (product *ProductManagement) UploadProduct(img *multipart.FileHeader, products *domain.ProductRequest) error {
	return nil
}

func (product *ProductManagement) GetAllSuppliers(ctx context.Context) ([]domain.Suppliers, error) {
	panic("not implemented")
}

func (product *ProductManagement) UploadNewsletter(ctx context.Context, news *domain.Newsletter) error {
	// TODO:
	return nil
}

func (product *ProductManagement) UploadHomeBanner(ctx context.Context, data *domain.HomeBanners) error {
	// TODO:
	return nil
}
