// Package comments implements comment
package comments

import (
	"context"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/pkg/infra/cache"
	"github.com/swclabs/swipex/pkg/infra/db"
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

// GetModelByProductID implements IComments.
func (comment *Comments) GetModelByProductID(ctx context.Context, ID int64) ([]model.Comment, error) {
	rows, err := comment.db.Query(ctx, getModelByProductID, ID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectRows[model.Comment](rows)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByProductID implements IComments.
func (comment *Comments) GetByProductID(ctx context.Context, ID int64) ([]entity.Comment, error) {
	rows, err := comment.db.Query(ctx, getByProductID, ID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectRows[entity.Comment](rows)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Insert implements ICommentRepository.
func (comment *Comments) Insert(ctx context.Context, cmt entity.Comment) (int64, error) {
	return comment.db.SafeWriteReturn(ctx, insert, cmt.Content, cmt.UserID, cmt.ProductID, cmt.InventoryID, cmt.StarID)
}

// GetByID implements ICommentRepository.
func (comment *Comments) GetByID(ctx context.Context, ID int64) (*entity.Comment, error) {
	row, err := comment.db.Query(ctx, getByID, ID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectRow[entity.Comment](row)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (comment *Comments) DeleteByID(ctx context.Context, ID int64) error {
	return comment.db.SafeWrite(ctx, deleteByID, ID)
}
