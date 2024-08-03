// Package collections implements collections
package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
)

type _cache struct {
	cache      cache.ICache
	collection ICollections
}

var _ ICollections = (*_cache)(nil)

func useCache(cache cache.ICache, collection ICollections) ICollections {
	return &_cache{
		cache:      cache,
		collection: collection,
	}
}

// AddCollection implements ICollections.
func (c *_cache) Create(ctx context.Context, banner entity.Collection) (int64, error) {
	return c.collection.Create(ctx, banner)
}

// AddHeadlineBanner implements ICollections.
func (c *_cache) AddHeadlineBanner(ctx context.Context, headline entity.Collection) error {
	return c.collection.AddHeadlineBanner(ctx, headline)
}

// GetMany implements ICollections.
func (c *_cache) GetMany(ctx context.Context, position string, limit int) ([]entity.Collection, error) {
	return c.collection.GetMany(ctx, position, limit)
}

// UploadCollectionImage implements ICollections.
func (c *_cache) UploadCollectionImage(ctx context.Context, collectionID string, url string) error {
	return c.collection.UploadCollectionImage(ctx, collectionID, url)
}
