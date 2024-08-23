package article

import (
	"context"
	"mime/multipart"
	"swclabs/swix/internal/core/domain/dtos"
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

	// UploadMessage uploads a new message.
	// ctx is the context to manage the request's lifecycle.
	// message contains the message details to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadMessage(ctx context.Context, message dtos.Message) error

	// GetMessage return a message.
	// ctx is the context to manage the request's lifecycle.
	// position contains the position of the message to be returns.
	// limit is the maximum number of message to retrieve.
	// Returns a message and an error if any issues occur during the upload process.
	GetMessage(ctx context.Context, position string, limit int) (*dtos.Message, error)

	// GetComment return a comment.
	// ctx is the context to manage the request's lifecycle.
	// position contains the position of the comment to be returns.
	// limit is the maximum number of comment to retrieve.
	// Returns a comment and an error if any issues occur during the upload process.
	GetComment(ctx context.Context, position string, limit int) (*dtos.Comment, error)

	// UploadComment uploads a new comment.
	// ctx is the context to manage the request's lifecycle.
	// comment contains the comment details to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadComment(ctx context.Context, comment dtos.Comment) error
}
