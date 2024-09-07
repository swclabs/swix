// Package comments implements comments
package comments

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

type ICommentRepository interface {
	// InsertComment inserts a new comment
	// ctx is the context to manage the request's lifecycle
	// cmt is the comment to be inserted
	// Returns the ID of the newly inserted comment and an error if any issues occur during the insertion process
	Insert(ctx context.Context, cmt entity.Comments) (int64, error)

	// GetByID retrieves a comment by its ID
	// ctx is the context to manage the request's lifecycle
	// id is the ID of the comment to be retrieved
	// Returns the comment and an error if any issues occur during the retrieval process
	GetByID(ctx context.Context, id int64) (*entity.Comments, error)

	// Update updates a comment
	// ctx is the context to manage the request's lifecycle
	// cmt is the comment to be updated
	// Returns an error if any issues occur during the update process
	Update(ctx context.Context, cmt entity.Comments) error

	// GetByProductID retrieves all comments for a product
	// ctx is the context to manage the request's lifecycle
	// productID is the ID of the product whose comments are to be retrieved
	// Returns the comments and an error if any issues occur during the retrieval process
	// GetByProductID(ctx context.Context, productID int64) ([]entity.Comments, error)

	// Delete deletes a comment
	// ctx is the context to manage the request's lifecycle
	// id is the ID of the comment to be deleted
	// Returns an error if any issues occur during the deletion process
	DeleteByID(ctx context.Context, id int64) error
}
