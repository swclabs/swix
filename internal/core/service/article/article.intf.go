package article

import (
	"context"
	"mime/multipart"

	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/model"
)

// IArticle : Module for managing article.
// Actor: Admin & Customer
type IArticle interface {
	UploadNews(ctx context.Context, newsDTO dtos.NewsDTO) (int64, error)

	UploadNewsImage(ctx context.Context, newsID int64, file *multipart.FileHeader) error

	GetNews(ctx context.Context, category string, limit int) (*dtos.News, error)

	// GetComment return a comment.
	// ctx is the context to manage the request's lifecycle.
	// position contains the position of the comment to be returns.
	// limit is the maximum number of comment to retrieve.
	// Returns a comment and an error if any issues occur during the upload process.
	GetComment(ctx context.Context, productID int64) ([]model.Comment, error)

	// UploadComment uploads a new comment.
	// ctx is the context to manage the request's lifecycle.
	// comment contains the comment details to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadComment(ctx context.Context, comment dtos.Comment) error
}
