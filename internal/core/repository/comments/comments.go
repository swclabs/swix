// Package comments implements comment
package comments

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
	"swclabs/swix/pkg/lib/errors"
)

var _ ICommentRepository = (*Comments)(nil)

func New(conn db.IDatabase) ICommentRepository {
	return &Comments{
		db: conn,
	}
}

func Init(conn db.IDatabase, cache cache.ICache) ICommentRepository {
	return useCache(cache, New(conn))
}

type Comments struct {
	db db.IDatabase
}

// Insert implements ICommentRepository.
func (comment *Comments) Insert(ctx context.Context, cmt entity.Comment) (int64, error) {
	id, err := comment.db.SafeWriteReturn(ctx, insertIntoComments,
		cmt.Level,
		cmt.Content,
		cmt.UserID,
		cmt.ProductID,
	)

	if err != nil {
		return -1, errors.Repository("write data", err)
	}
	return id, nil
}
