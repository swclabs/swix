package service

import (
	"context"
	"mime/multipart"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
	"swclabs/swipe-api/pkg/cloud"
)

type ProductManagement struct {
	Category  domain.ICategoriesRepository
	product   domain.IProductRepository
	suppliers domain.ISuppliersRepository
}

func NewProductManagement() domain.IProductManagementService {
	return &ProductManagement{
		Category:  repo.NewCategories(),
		product:   repo.NewProducts(),
		suppliers: repo.NewSuppliers(),
	}
}

func (product *ProductManagement) InsertCategory(ctx context.Context, ctg *domain.Categories) error {
	// call repository layer
	return product.Category.Insert(ctx, ctg)
}

func (product *ProductManagement) UploadImage(Id string, fileHeader *multipart.FileHeader) error {
	return nil
}

func (product *ProductManagement) UploadProduct(ctx context.Context, fileHeader *multipart.FileHeader, products domain.ProductRequest) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UpdateImages(cloud.Connection(), file)
	if err != nil {
		return err
	}
	var prd = domain.Products{
		Image:       resp.SecureURL,
		Price:       products.Price,
		Description: products.Description,
		Name:        products.Name,
		SupplierID:  products.SupplierID,
		CategoryID:  products.CategoryID,
		Available:   products.Available,
	}
	return product.product.Insert(ctx, &prd)
}

func (product *ProductManagement) GetSuppliersLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	return product.suppliers.GetLimit(ctx, limit)
}

func (product *ProductManagement) InsertSuppliers(ctx context.Context, supplierReq domain.SuppliersRequest) error {
	supplier := domain.Suppliers{
		Name:        supplierReq.Name,
		Email:       supplierReq.Email,
		PhoneNumber: supplierReq.PhoneNumber,
	}
	addr := domain.Addresses{
		City: supplierReq.City,
		Ward: supplierReq.Ward,
		District: supplierReq.District,
		Street: supplierReq.Street,
	}
	return product.suppliers.Insert(ctx, supplier, addr)
}

func (product *ProductManagement) UploadNewsletter(ctx context.Context, news domain.Newsletter, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UpdateImages(cloud.Connection(), file)
	if err != nil {
		return err
	}
	news.Image = resp.SecureURL
	return repo.NewNewsletter().Insert(context.Background(), news)
}

func (product *ProductManagement) UploadHomeBanner(ctx context.Context, data *domain.HomeBanners) error {
	// TODO:
	return nil
}
