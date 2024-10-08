package products

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/enum"
	"swclabs/swix/internal/core/repos/categories"
	"swclabs/swix/internal/core/repos/inventories"
	"swclabs/swix/internal/core/repos/products"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/infra/db"
	"swclabs/swix/pkg/lib/errors"
	"swclabs/swix/pkg/utils"

	"github.com/shopspring/decimal"
)

var _ IProducts = (*Products)(nil)
var _ = app.Service(New)

// New creates a new ProductService object
func New(
	blob blob.IBlobStorage,
	products products.IProducts,
	inventory inventories.IInventories,
	category categories.ICategories,
) IProducts {
	return &Products{
		Blob:      blob,
		Products:  products,
		Inventory: inventory,
		Category:  category,
	}
}

// Products struct for product service
type Products struct {
	Blob      blob.IBlobStorage
	Products  products.IProducts
	Inventory inventories.IInventories
	Category  categories.ICategories
}

// SearchDetails implements IProducts.
func (s *Products) SearchDetails(ctx context.Context, keyword string) ([]dtos.ProductDetail, error) {
	products, err := s.Products.Search(ctx, keyword)
	if err != nil {
		return nil, errors.Service("keyword error", err)
	}
	var details []dtos.ProductDetail
	for _, product := range products {
		detail, err := s.ProductDetail(ctx, product.ID)
		if err != nil {
			return nil, err
		}
		details = append(details, *detail)
	}
	return details, nil
}

// AccessoryDetail implements IProductService.
func (s *Products) AccessoryDetail(ctx context.Context, productID int64) (*dtos.AccessoryDetail, error) {
	product, err := s.Products.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	category, err := s.Category.GetByID(ctx, product.CategoryID)
	if err != nil {
		return nil, err
	}
	var types enum.Category
	if err := types.Load(category.Name); err != nil {
		return nil, err
	}
	if types&enum.Accessory == 0 {
		return &dtos.AccessoryDetail{}, nil
	}
	var detail = dtos.AccessoryDetail{
		Name:     product.Name,
		Price:    product.Price,
		Status:   product.Status,
		Image:    strings.Split(product.Image, ","),
		Category: types.String(),
	}
	return &detail, nil
}

// ViewDataOf implements IProductService.
func (s *Products) ViewDataOf(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductView, error) {
	products, err := s.Products.GetByCategory(ctx, types, offset)
	if err != nil {
		return nil, err
	}
	var productView []dtos.ProductView
	for _, p := range products {
		_view := dtos.ProductView{
			ID:       p.ID,
			Price:    p.Price,
			Desc:     p.Description,
			Name:     p.Name,
			Image:    p.Image,
			Category: p.CategoryName,
		}
		if p.Specs != "" && types&enum.Storage != 0 {
			var specs dtos.ProductSpecs
			if err := json.Unmarshal([]byte(p.Specs), &specs); err != nil {
				return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
			}
			_view.Specs = specs
		}
		productView = append(productView, _view)
	}
	return productView, nil
}

// GetInvByID implements IProductService.
func (s *Products) GetInvByID(ctx context.Context, inventoryID int64) (*dtos.Inventory, error) {
	item, err := s.Inventory.GetByID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	product, err := s.Products.GetByID(ctx, item.ProductID)
	if err != nil {
		return nil, err
	}
	category, _ := s.Category.GetByID(ctx, product.CategoryID)
	var specs dtos.Specs
	if err := json.Unmarshal([]byte(item.Specs), &specs); err != nil {
		return nil, err
	}
	var (
		result = dtos.Inventory{
			ID:           item.ID,
			ProductName:  product.Name,
			ProductID:    item.ProductID,
			Price:        item.Price.String(),
			Available:    strconv.Itoa(int(item.Available)),
			CurrencyCode: item.CurrencyCode,
			Status:       item.Status,
			Color:        item.Color,
			ColorImg:     item.ColorImg,
			Image:        strings.Split(item.Image, ","),
			Category:     category.Name,
			Specs:        []dtos.Specs{specs},
		}
	)
	return &result, nil

}

// ProductDetail implements IProductService.
func (s *Products) ProductDetail(ctx context.Context, productID int64) (*dtos.ProductDetail, error) {
	var (
		productSpecs dtos.ProductSpecs
		details      dtos.ProductDetail
	)
	colors, err := s.Inventory.GetColor(ctx, productID)
	if err != nil {
		return nil, err
	}
	product, err := s.Products.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(product.Specs), &productSpecs); err != nil {
		return nil, err
	}
	details.Name = product.Name
	details.Screen = productSpecs.Screen
	details.Display = productSpecs.Display
	details.Image = strings.Split(product.Image, ",")

	for _, color := range colors {
		items, err := s.Inventory.GetByColor(ctx, productID, color.Color)
		if err != nil {
			return nil, err
		}
		detailsColor := dtos.DetailColor{
			Name:    color.Color,
			Img:     items[0].ColorImg,
			Product: strings.Split(items[0].Image, ","),
		}
		for _, item := range items {
			var spec dtos.SpecsItem
			if err := json.Unmarshal([]byte(item.Specs), &spec); err != nil {
				return nil, err
			}
			spec.Price = item.Price.String()
			detailsColor.Specs = append(detailsColor.Specs, spec)
		}
		details.Color = append(details.Color, detailsColor)
	}
	return &details, nil
}

// UpdateInv implements IProductService.
func (s *Products) UpdateInv(ctx context.Context, inventory dtos.InvUpdate) error {
	pid, err := strconv.Atoi(inventory.ProductID)
	if err != nil {
		pid = -1
	}
	price, err := decimal.NewFromString(inventory.Price)
	if err != nil {
		price = decimal.NewFromInt(-1)
	}
	avai, err := strconv.ParseInt(inventory.Available, 10, 64)
	if err != nil {
		avai = -1
	}
	invID, _ := strconv.ParseInt(inventory.ID, 10, 64)
	return s.Inventory.Update(ctx, entity.Inventories{
		Price:        price,
		ID:           invID,
		Available:    avai,
		ProductID:    int64(pid),
		Status:       inventory.Status,
		CurrencyCode: inventory.CurrencyCode,
	})
}

// UploadInvImage implements IProductService.
func (s *Products) UploadInvImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := s.Blob.UploadImages(file)
		if err == nil {
			if err = s.Inventory.UploadImage(ctx, ID, resp.SecureURL); err == nil {
				if err = file.Close(); err != nil {
					return err
				}
			}
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteInvByID implements IProductService.
func (s *Products) DeleteInvByID(ctx context.Context, inventoryID int64) error {
	return s.Inventory.DeleteByID(ctx, inventoryID)
}

// GetAllInv implements IProductService.
func (s *Products) GetAllInv(ctx context.Context, page int, limit int) (*dtos.InvItems, error) {

	products, err := s.Products.GetLimit(ctx, limit)
	if err != nil {
		return nil, err
	}
	var invItems dtos.InvItems
	for _, product := range products {
		colors, err := s.Inventory.GetColor(ctx, product.ID)
		if err != nil {
			return nil, err
		}
		category, err := s.Category.GetByID(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}
		for _, color := range colors {
			items, err := s.Inventory.GetByColor(ctx, product.ID, color.Color)
			if err != nil {
				return nil, err
			}
			_item := dtos.Inventory{
				ID:           items[0].ID,
				ProductName:  product.Name,
				ProductID:    items[0].ProductID,
				Price:        items[0].Price.String(),
				Available:    strconv.Itoa(int(items[0].Available)),
				CurrencyCode: items[0].CurrencyCode,
				Status:       items[0].Status,
				Color:        items[0].Color,
				ColorImg:     items[0].ColorImg,
				Image:        strings.Split(items[0].Image, ","),
				Category:     category.Name,
				Specs:        nil,
			}
			for _, item := range items {
				switch item.Status {
				case "active":
					invItems.Header.Active++
				case "draft":
					invItems.Header.Draft++
				case "archived":
					invItems.Header.Active++
				}
				invItems.Header.All++
				var specs dtos.Specs
				if err := json.Unmarshal([]byte(item.Specs), &specs); err != nil {
					return nil, err
				}
				_item.Specs = append(_item.Specs, specs)
			}
			invItems.Stock = append(invItems.Stock, _item)
		}
	}

	invItems.Page = page
	invItems.Limit = limit

	return &invItems, nil
}

// GetInv implements IProductService.
func (s *Products) GetInv(ctx context.Context, productID int64) ([]entity.Inventories, error) {
	return s.Inventory.GetByProductID(ctx, productID)
}

// Search implements IProductService.
func (s *Products) Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error) {
	_products, err := s.Products.Search(ctx, keyword)
	if err != nil {
		return nil, errors.Service("keyword error", err)
	}
	var productSchema = []dtos.ProductResponse{}
	for _, p := range _products {
		category, err := s.Category.GetByID(ctx, p.CategoryID)
		if err != nil {
			return nil, err
		}
		productSchema = append(productSchema, dtos.ProductResponse{
			ID:          p.ID,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
			Status:      p.Status,
			Image:       strings.Split(p.Image, ","),
			Created:     utils.HanoiTimezone(p.Created),
			Category:    category.Name,
		})
	}
	return productSchema, nil
}

// UpdateProductInfo implements IProductService.
func (s *Products) UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error {
	if product.CategoryID != 0 {
		_category, err := s.Category.GetByID(ctx, product.CategoryID)
		if err != nil {
			return fmt.Errorf("category not found %v", err)
		}
		var types enum.Category
		if err := types.Load(_category.Name); err != nil {
			return fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
		}
	}
	var (
		sampleSpec, _ = json.Marshal(dtos.ProductSpecs{})
		spec, _       = json.Marshal(product.Specs)
		sspec         = ""
	)
	if string(sampleSpec) != string(spec) {
		sspec = string(spec)
	}
	_product := entity.Products{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		SupplierID:  product.SupplierID,
		CategoryID:  product.CategoryID,
		Status:      product.Status,
		Specs:       sspec,
	}
	return s.Products.Update(ctx, _product)

}

// UploadProductImage implements IProductService.
func (s *Products) UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := s.Blob.UploadImages(file)
		if err != nil {
			return err
		}
		if err := s.Products.UploadNewImage(ctx, resp.SecureURL, ID); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}

// CreateProduct implements IProductService.
func (s *Products) CreateProduct(ctx context.Context, products dtos.Product) (int64, error) {
	_category, err := s.Category.GetByID(ctx, products.CategoryID)
	if err != nil {
		return -1, fmt.Errorf("category not found %v", err)
	}

	var types enum.Category
	if err := types.Load(_category.Name); err != nil {
		return -1, fmt.Errorf("category invalid %v", err)
	}
	var prd = entity.Products{
		Price:       products.Price,
		Description: products.Description,
		Name:        products.Name,
		SupplierID:  products.SupplierID,
		CategoryID:  products.CategoryID,
		Status:      products.Status,
		Specs:       "{}",
	}
	if products.Specs != nil {
		var specs, ok = products.Specs.(dtos.ProductSpecs)
		if !ok {
			return -1, fmt.Errorf("[code: %d] invalid specifications", http.StatusBadRequest)
		}
		specsByte, _ := json.Marshal(specs)
		prd.Specs = string(specsByte)
	}
	return s.Products.Insert(ctx, prd)
}

// DelProductByID implements IProductService.
func (s *Products) DelProductByID(ctx context.Context, productID int64) error {
	return s.Products.DeleteByID(ctx, productID)
}

// InsertInv implements IProductService.
func (s *Products) InsertInv(ctx context.Context, product dtos.Inventory) error {
	var (
		price, _  = decimal.NewFromString(product.Price)
		avai, _   = strconv.Atoi(product.Available)
		inventory = entity.Inventories{
			Color:        product.Color,
			ColorImg:     product.ColorImg,
			Image:        strings.Join(product.Image, ","),
			ProductID:    product.ProductID,
			Price:        price,
			Available:    int64(avai),
			CurrencyCode: product.CurrencyCode,
			Status:       "active",
		}
	)
	tx, err := db.NewTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		invRepo = inventories.New(tx)
	)
	for _, spec := range product.Specs {
		specs, _ := json.Marshal(spec)
		inventory.Specs = string(specs)
		_, err := invRepo.InsertProduct(ctx, inventory)
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				return errTx
			}
			return err
		}
	}
	return tx.Commit(ctx)
}

// GetProductsLimit implements IProductService.
func (s *Products) GetProductsLimit(ctx context.Context, limit int) ([]dtos.ProductResponse, error) {
	products, err := s.Products.GetLimit(ctx, limit)
	if err != nil {
		return nil, err
	}
	var productResponse []dtos.ProductResponse
	for _, p := range products {
		var (
			product = dtos.ProductResponse{
				ID:          p.ID,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
				Status:      p.Status,
				Created:     utils.HanoiTimezone(p.Created),
				Image:       strings.Split(p.Image, ","),
			}
			types enum.Category
		)
		category, err := s.Category.GetByID(ctx, p.CategoryID)
		if err != nil {
			return nil, err
		}
		if err := types.Load(category.Name); err != nil {
			return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
		}
		product.Category = category.Name
		productResponse = append(productResponse, product)
	}
	return productResponse, nil
}
