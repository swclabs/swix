package service

import (
	"context"
	"mime/multipart"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
	"swclabs/swipe-api/pkg/cloud"
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
	return product.Category.Insert(ctx, ctg)
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

func (product *ProductManagement) UploadNewsletter(ctx context.Context, news domain.Newsletter, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UpdateImages(cloud.Connection(), file)
	if err != nil {
		return nil
	}
	news.Image = resp.SecureURL
	return repo.NewNewsletter().Insert(context.Background(), news)
}

func (product *ProductManagement) UploadHomeBanner(ctx context.Context, data *domain.HomeBanners) error {
	// TODO:
	return nil
}
