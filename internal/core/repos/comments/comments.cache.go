package comments

import (
	"context"
	//"crypto"
	"fmt"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/lib/crypto"
)

var _ ICommentRepository = (*_cache)(nil)

var (
	_                 ICommentRepository = (*_cache)(nil)
	keyGetByID                           = "ICommentRepository:GetByID:%d"
	keyGetByProductID                    = "ICommentRepository:GetByProductID:%d"
)

func useCache(cache cache.ICache, comment ICommentRepository) ICommentRepository {
	return &_cache{
		comments: comment,
		cache:    cache,
	}
}

type _cache struct {
	cache    cache.ICache
	comments ICommentRepository
}

// Insert implements ICommentRepository.
func (c *_cache) Insert(ctx context.Context, cmt entity.Comments) (int64, error) {
	return c.comments.Insert(ctx, cmt)
}

// GetByID implements ICommentRepository.
func (c *_cache) GetByID(ctx context.Context, commentID int64) (*entity.Comments, error) {
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, commentID))

	result, err := cache.Get[entity.Comments](ctx, c.cache, key)
	if err != nil {
		// return nil, err
		result, err = c.comments.GetByID(ctx, commentID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[entity.Comments](ctx, c.cache, key, *result); err != nil {
			return result, err
		}
	}
	return result, nil
}

// Update implements ICommentRepository.
func (c *_cache) Update(ctx context.Context, cmt entity.Comments) error {
	return c.comments.Update(ctx, cmt)
}

// GetByProductID implements ICommentRepository.
func (c *_cache) GetByProductID(ctx context.Context, ID int64) ([]entity.Comments, error) {
	// return c.comments.GetByProductID(ctx, productID)
	key := crypto.HashOf(fmt.Sprintf(keyGetByProductID, ID))

	result, err := cache.GetSlice[entity.Comments](ctx, c.cache, key)
	if err != nil {
		// return nil, err
		result, err = c.comments.GetByProductID(ctx, ID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[[]entity.Comments](ctx, c.cache, key, result); err != nil {
			return result, err
		}
	}
	return result, nil
}

// Delete implements ICommentRepository.
func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.comments.DeleteByID(ctx, ID)
}
