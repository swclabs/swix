package products

import (
	"context"
	"encoding/json"
	"errors"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/repository/suppliers"
	"swclabs/swipecore/internal/core/repository/warehouse"
	"swclabs/swipecore/pkg/blob"
)

type ProductService struct {
	Categories categories.ICategoriesRepository
	Products   products.IProductRepository
	Suppliers  suppliers.ISuppliersRepository
	Warehouse  warehouse.IWarehouseRepository
}

// New creates a new ProductService instance
func New() IProductService {
	return &ProductService{
		Categories: categories.New(),
		Products:   products.New(),
		Suppliers:  suppliers.New(),
		Warehouse:  warehouse.New(),
	}
}

// GetProductsInWarehouse implements domain.IProductService.
func (s *ProductService) GetProductsInWarehouse(ctx context.Context, productID, ram, ssd, color string) (*domain.WarehouseType, error) {
	warehouse, err := s.Warehouse.GetProducts(ctx, productID, ram, ssd, color)
	if err != nil {
		return nil, err
	}
	var warehouseRes = domain.WarehouseType{
		Id: warehouse.Id,
		WarehouseStructure: domain.WarehouseStructure{
			ProductID: warehouse.Id,
			Price:     warehouse.Price,
			Model:     warehouse.Model,
			Available: warehouse.Available,
		},
	}
	if err := json.Unmarshal([]byte(warehouse.Specs), &warehouseRes.Specs); err != nil {
		return &warehouseRes, nil // don't find anything, just return empty object
	}
	if warehouseRes.Available == "" {
		warehouseRes.Available = "0"
		return &warehouseRes, nil
	}
	return &warehouseRes, nil
}

// InsertIntoWarehouse implements domain.IProductService.
func (s *ProductService) InsertIntoWarehouse(ctx context.Context, product domain.WarehouseStructure) error {
	return s.Warehouse.InsertProduct(ctx, product)
}

func (s *ProductService) GetCategoriesLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	return s.Categories.GetLimit(ctx, limit)
}

func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]domain.ProductRes, error) {
	return s.Products.GetLimit(ctx, limit)
}

func (s *ProductService) InsertCategory(ctx context.Context, ctg *domain.Categories) error {
	return s.Categories.Insert(ctx, ctg)
}

func (s *ProductService) GetSuppliersLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	return s.Suppliers.GetLimit(ctx, limit)
}

func (s *ProductService) UploadProductImage(ctx context.Context, Id int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return errors.New("missing image file")
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := blob.UploadImages(blob.Connection(), file)
		if err != nil {
			return err
		}
		if err := s.Products.UploadNewImage(ctx, resp.SecureURL, Id); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProductService) UploadProduct(ctx context.Context, products domain.ProductReq) (int64, error) {
	specs, err := json.Marshal(domain.Specs{
		Screen:  products.Screen,
		Display: products.Display,
		SSD:     products.SSD,
		RAM:     products.RAM,
	})
	if err != nil {
		return -1, err
	}
	var prd = domain.Products{
		Price:       products.Price,
		Description: products.Description,
		Name:        products.Name,
		SupplierID:  products.SupplierID,
		CategoryID:  products.CategoryID,
		Status:      products.Status,
		Spec:        string(specs),
	}
	return s.Products.Insert(ctx, &prd)
}

func (s *ProductService) InsertSuppliers(ctx context.Context, supplierReq domain.SuppliersReq) error {
	supplier := domain.Suppliers{
		Name:  supplierReq.Name,
		Email: supplierReq.Email,
	}
	addr := domain.Addresses{
		City:     supplierReq.City,
		Ward:     supplierReq.Ward,
		District: supplierReq.District,
		Street:   supplierReq.Street,
	}
	return s.Suppliers.Insert(ctx, supplier, addr)
}
