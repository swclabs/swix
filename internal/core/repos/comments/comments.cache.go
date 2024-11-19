package comments

import (
	"context"
	//"crypto"
	"fmt"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/pkg/infra/cache"
	"github.com/swclabs/swipex/pkg/lib/crypto"
)

var _ IComments = (*_cache)(nil)

var (
	_          IComments = (*_cache)(nil)
	keyGetByID           = "ICommentRepository:GetByID:%d"
	// keyGetByProductID           = "ICommentRepository:GetByProductID:%d"
)

func useCache(cache cache.ICache, comment IComments) IComments {
	return &_cache{
		comments: comment,
		cache:    cache,
	}
}

type _cache struct {
	cache    cache.ICache
	comments IComments
}

// GetModelByProductID implements IComments.
func (c *_cache) GetModelByProductID(ctx context.Context, ID int64) ([]model.Comment, error) {
	return c.comments.GetModelByProductID(ctx, ID)
}

// GetByProductID implements IComments.
func (c *_cache) GetByProductID(ctx context.Context, ID int64) ([]entity.Comment, error) {
	return c.comments.GetByProductID(ctx, ID)
}

// Insert implements ICommentRepository.
func (c *_cache) Insert(ctx context.Context, cmt entity.Comment) (int64, error) {
	return c.comments.Insert(ctx, cmt)
}

// GetByID implements ICommentRepository.
func (c *_cache) GetByID(ctx context.Context, commentID int64) (*entity.Comment, error) {
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, commentID))

	result, err := cache.Get[entity.Comment](ctx, c.cache, key)
	if err != nil {
		// return nil, err
		result, err = c.comments.GetByID(ctx, commentID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[entity.Comment](ctx, c.cache, key, *result); err != nil {
			return result, err
		}
	}
	return result, nil
}

// Delete implements ICommentRepository.
func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.comments.DeleteByID(ctx, ID)
}
