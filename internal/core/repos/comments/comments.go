// Package comments implements comment
package comments

import (
	"context"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/cache"
	"github.com/swclabs/swipex/pkg/infra/db"
	"github.com/swclabs/swipex/pkg/lib/errors"
)

var _ IComments = (*Comments)(nil)
var _ = app.Repos(Init)

func New(conn db.IDatabase) IComments {
	return &Comments{
		db: conn,
	}
}

func Init(conn db.IDatabase, cache cache.ICache) IComments {
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
		cmt.ParentID,
	)

	if err != nil {
		return -1, errors.Repository("write data", err)
	}
	return id, nil
}

// GetByID implements ICommentRepository.
func (comment *Comments) GetByID(ctx context.Context, ID int64) (*entity.Comment, error) {
	row, err := comment.db.Query(ctx, selectByID, ID)

	if err != nil {
		return nil, errors.Repository("query", err)
	}

	_comment, err := db.CollectRow[entity.Comment](row)
	if err != nil {
		return nil, err
	}
	return &_comment, nil
}

// Update implements ICommentRepository.
func (comment *Comments) Update(ctx context.Context, cmt entity.Comment) error {
	return comment.db.SafeWrite(ctx, updateComments,
		cmt.ID,
		cmt.Content,
		cmt.Level,
		cmt.ProductID,
		cmt.UserID,
	)
}

// GetByProductID implements ICommentRepository.
func (comment *Comments) GetByProductID(ctx context.Context, productID int64) ([]entity.Comment, error) {
	rows, err := comment.db.Query(ctx, selectCommentsByProductID, productID)
	if err != nil {
		return nil, err
		// return nil, errors.Repository("500", err)
	}
	comments, err := db.CollectRows[entity.Comment](rows)
	if err != nil {
		// return nil, errors.Repository("500", err)
		return nil, err
	}
	return comments, nil
}

func (comment *Comments) DeleteByID(ctx context.Context, ID int64) error {
	return errors.Repository("safely write data",
		comment.db.SafeWrite(ctx, deleteByID, ID),
	)
}
