package service

import (
	"mime/multipart"

	"github.com/swclabs/swipe-api/internal/core/domain"
	"github.com/swclabs/swipe-api/internal/core/repo"
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
	return product.Category.New(ctg)
}

func (product *ProductManagement) UploadImage(Id string, fileHeader *multipart.FileHeader) error {
	return nil
}
