package service

import (
	"mime/multipart"

	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/repo"
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

func (product *ProductManagement) UploadImage(Id int64, fileHeader *multipart.FileHeader) error {
	return nil
}
