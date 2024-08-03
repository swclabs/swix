package article

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain/dtos"
)

// IArticle : Module for managing article.
// Actor: Admin & Customer
type IArticle interface {
	// UploadArticle uploads a new collection.
	// ctx is the context to manage the request's lifecycle.
	// banner contains the collection details to be uploaded.
	// Returns id of collection was uploaded and error if any issues occur during the upload process.
	UploadArticle(ctx context.Context, banner dtos.UploadArticle) (int64, error)

	// UploadCollectionsImage uploads a new image of collection.
	// ctx is the context to manage the request's lifecycle.
	// cardBannerID contains the id of collection to be uploaded.
	// fileHeader is  the header of the file to be uploaded
	// Returns an error if any issues occur during the upload process.
	UploadCollectionsImage(ctx context.Context, cardBannerID string, fileHeader *multipart.FileHeader) error

	// GetCarousels return a slices of carousel.
	// ctx is the context to manage the request's lifecycle.
	// cardBannerID contains the id of collection to be returns.
	// limit is the maximum number of Collection to retrieve.
	// Returns an error if any issues occur during the upload process.
	GetCarousels(ctx context.Context, position string, limit int) (*dtos.Article, error)

	// UploadHeadlineBanner uploads a new headline banner.
	// ctx is the context to manage the request's lifecycle.
	// banner contains the headline banner details to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadHeadlineBanner(ctx context.Context, banner dtos.HeadlineBanner) error

	// SliceOfHeadlineBanner return a slices of headline banner.
	// ctx is the context to manage the request's lifecycle.
	// position contains the position of headline banner to be returns.
	// limit is the maximum number of headline banner to retrieve.
	// Returns an error if any issues occur during the upload process.
	SliceOfHeadlineBanner(ctx context.Context, position string, limit int) (*dtos.HeadlineBanners, error)
}
