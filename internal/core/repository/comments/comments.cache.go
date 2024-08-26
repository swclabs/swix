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
func (c *_cache) Insert(ctx context.Context, cmt entity.Comment) (int64, error) {
	return c.comments.Insert(ctx, cmt)
}
