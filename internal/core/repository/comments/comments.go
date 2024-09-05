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
func (comment *Comments) Insert(ctx context.Context, cmt entity.Comments) (int64, error) {
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

// GetByID implements ICommentRepository.
func (comment *Comments) GetByID(ctx context.Context, ID int64) (*entity.Comments, error) {
	row, err := comment.db.Query(ctx, selectCommentByID, ID)

	if err != nil {
		return nil, err
	}
	result, err := db.CollectOneRow[entity.Comments](row)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update implements ICommentRepository.
func (comment *Comments) Update(ctx context.Context, cmt entity.Comments) error {
	return comment.db.SafeWrite(ctx, updateComments,
		cmt.ID,
		cmt.Content,
		cmt.Level,
		cmt.ProductID,
		cmt.UserID,
	)
}

// GetCommentsByProductID implements ICommentRepository.
func (comment *Comments) GetByProductID(ctx context.Context, productID int64) ([]entity.Comments, error) {
	rows, err := comment.db.Query(ctx, selectCommentsByProductID, productID)
	if err != nil {
		return nil, err
	}
	comments, err := db.CollectRows[entity.Comments](rows)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
