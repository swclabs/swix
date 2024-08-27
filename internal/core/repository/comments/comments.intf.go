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
	Insert(ctx context.Context, cmt entity.Comment) (int64, error)

	// GetByID retrieves a comment by its ID
	// ctx is the context to manage the request's lifecycle
	// id is the ID of the comment to be retrieved
	// Returns the comment and an error if any issues occur during the retrieval process
	GetByID(ctx context.Context, id int64) (*entity.Comment, error)
}
