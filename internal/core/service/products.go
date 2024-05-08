package service

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository"
	"swclabs/swipecore/pkg/cloud"
)

type ProductService struct {
	Categories domain.ICategoriesRepository
	Products   domain.IProductRepository
	Suppliers  domain.ISuppliersRepository
	Warehouse  domain.IWarehouseRepository
}

// NewProductService creates a new ProductService instance
func NewProductService() domain.IProductService {
	return &ProductService{
		Categories: repository.NewCategories(),
		Products:   repository.NewProducts(),
		Suppliers:  repository.NewSuppliers(),
		Warehouse:  repository.NewWarehouse(),
	}
}

// GetProductsInWarehouse implements domain.IProductService.
func (s *ProductService) GetProductsInWarehouse(ctx context.Context, productID, ram, ssd string) (*domain.Warehouse, error) {
	return s.Warehouse.GetProducts(ctx, productID, ram, ssd)
}

// InsertIntoWarehouse implements domain.IProductService.
func (s *ProductService) InsertIntoWarehouse(ctx context.Context, product domain.Warehouse) error {
	return s.Warehouse.InsertProduct(ctx, product)
}

func (s *ProductService) GetCategoriesLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	return s.Categories.GetLimit(ctx, limit)
}

func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]domain.Products, error) {
	return s.Products.GetLitmit(ctx, limit)
}

func (s *ProductService) InsertCategory(ctx context.Context, ctg *domain.Categories) error {
	return s.Categories.Insert(ctx, ctg)
}

func (s *ProductService) GetSuppliersLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	return s.Suppliers.GetLimit(ctx, limit)
}

func (s *ProductService) UploadProductImage(ctx context.Context, Id int, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UploadImages(cloud.Connection(), file)
	if err != nil {
		return err
	}
	return s.Products.UploadNewImage(ctx, resp.SecureURL, Id)
}

func (s *ProductService) UploadProduct(ctx context.Context, fileHeader *multipart.FileHeader, products domain.ProductRequest) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UploadImages(cloud.Connection(), file)
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
		Status:      products.Status,
		Spec:        products.Spec,
	}
	return s.Products.Insert(ctx, &prd)
}

func (s *ProductService) InsertSuppliers(ctx context.Context, supplierReq domain.SuppliersRequest) error {
	supplier := domain.Suppliers{
		Name:        supplierReq.Name,
		Email:       supplierReq.Email,
		PhoneNumber: supplierReq.PhoneNumber,
	}
	addr := domain.Addresses{
		City:     supplierReq.City,
		Ward:     supplierReq.Ward,
		District: supplierReq.District,
		Street:   supplierReq.Street,
	}
	return s.Suppliers.Insert(ctx, supplier, addr)
}
