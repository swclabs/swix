package comments

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
)

var _ ICommentRepository = (*_cache)(nil)

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
func (c *_cache) GetByID(ctx context.Context, productID int64) (*entity.Comments, error) {
	return c.comments.GetByID(ctx, productID)
}

// Update implements ICommentRepository.
func (c *_cache) Update(ctx context.Context, cmt entity.Comments) error {
	return c.comments.Update(ctx, cmt)
}

// GetByProductID implements ICommentRepository.
// func (c *_cache) GetByProductID(ctx context.Context, productID int64) ([]entity.Comments, error) {
// 	return c.comments.GetByProductID(ctx, productID)
// }

// Delete implements ICommentRepository.
func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.comments.DeleteByID(ctx, ID)
}
