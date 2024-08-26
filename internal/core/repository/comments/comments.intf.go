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
}
