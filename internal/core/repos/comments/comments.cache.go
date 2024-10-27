package comments

import (
	"context"
	//"crypto"
	"fmt"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/cache"
	"swclabs/swipex/pkg/lib/crypto"
)

var _ IComments = (*_cache)(nil)

var (
	_                 IComments = (*_cache)(nil)
	keyGetByID                  = "ICommentRepository:GetByID:%d"
	keyGetByProductID           = "ICommentRepository:GetByProductID:%d"
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

// Update implements ICommentRepository.
func (c *_cache) Update(ctx context.Context, cmt entity.Comment) error {
	return c.comments.Update(ctx, cmt)
}

// GetByProductID implements ICommentRepository.
func (c *_cache) GetByProductID(ctx context.Context, ID int64) ([]entity.Comment, error) {
	// return c.comments.GetByProductID(ctx, productID)
	key := crypto.HashOf(fmt.Sprintf(keyGetByProductID, ID))

	result, err := cache.GetSlice[entity.Comment](ctx, c.cache, key)
	if err != nil {
		// return nil, err
		result, err = c.comments.GetByProductID(ctx, ID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[[]entity.Comment](ctx, c.cache, key, result); err != nil {
			return result, err
		}
	}
	return result, nil
}

// Delete implements ICommentRepository.
func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.comments.DeleteByID(ctx, ID)
}
