package products

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/enum"
	"github.com/swclabs/swipex/internal/core/repos/products"
	"github.com/swclabs/swipex/internal/core/repos/stars"
	"github.com/swclabs/swipex/pkg/infra/db"

	"github.com/shopspring/decimal"
)

// Rating implements IProducts.
func (p *Products) Rating(ctx context.Context, userID, productID int64, rating int) error {
	tx, err := db.NewTx(ctx)
	if err != nil {
		return err
	}

	var (
		star    = stars.New(tx)
		product = products.New(tx)
	)

	if _, err := star.Save(ctx, entity.Star{UserID: userID, ProductID: productID}); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	if err := product.Rating(ctx, productID, rating); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}

// UploadItemColorImage implements IProducts.
func (p *Products) UploadItemColorImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
	}

	file, err := fileHeader[0].Open()
	if err != nil {
		return err
	}

	resp, err := p.Blob.UploadImages(file)
	if err == nil {
		if err = p.Inventory.UploadColorImage(ctx, ID, resp.SecureURL); err == nil {
			if err = file.Close(); err != nil {
				return err
			}
		}
	}

	if err != nil {
		return err
	}
	return nil
}

// UploadProductShopImage implements IProducts.
func (p *Products) UploadProductShopImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := p.Blob.UploadImages(file)
		if err != nil {
			return err
		}
		if err := p.Products.UploadShopImage(ctx, resp.SecureURL, ID); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}

// UpdateItem implements IProductService.
func (p *Products) UpdateItem(ctx context.Context, inventory dtos.InvUpdate) error {
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
	return p.Inventory.Update(ctx, entity.Inventory{
		Price:        price,
		ID:           invID,
		Available:    avai,
		ProductID:    int64(pid),
		Status:       inventory.Status,
		CurrencyCode: inventory.CurrencyCode,
	})
}

// UploadItemImage implements IProductService.
func (p *Products) UploadItemImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := p.Blob.UploadImages(file)
		if err == nil {
			if err = p.Inventory.UploadImage(ctx, ID, resp.SecureURL); err == nil {
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

// UpdateProductInfo implements IProductService.
func (p *Products) UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error {
	if product.CategoryID != 0 {
		_category, err := p.Category.GetByID(ctx, product.CategoryID)
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
	_product := entity.Product{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		SupplierID:  product.SupplierID,
		CategoryID:  product.CategoryID,
		Status:      product.Status,
		Specs:       sspec,
	}
	return p.Products.Update(ctx, _product)

}

// UploadProductImage implements IProductService.
func (p *Products) UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := p.Blob.UploadImages(file)
		if err != nil {
			return err
		}
		if err := p.Products.UploadNewImage(ctx, resp.SecureURL, ID); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}
