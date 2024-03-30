package service

import (
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

func (product *ProductManagement) InsertCategory(ctg *domain.Categories) error {
	// call repository layer
	return product.Category.New(ctg)
}

func (product *ProductManagement) UploadImage(Id string, fileHeader *multipart.FileHeader) error {
	return nil
}

func (product *ProductManagement) UploadProduct(img *multipart.FileHeader, products *domain.ProductRequest) error {
	return nil
}
