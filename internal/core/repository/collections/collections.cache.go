// Package collections implements collections
package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
)

type cache struct {
	collection ICollections
}

var _ ICollections = (*cache)(nil)

func useCache(collection ICollections) ICollections {
	return &cache{collection: collection}
}

// AddCollection implements ICollections.
func (c *cache) AddCollection(ctx context.Context, banner entity.Collection) (int64, error) {
	return c.collection.AddCollection(ctx, banner)
}

// AddHeadlineBanner implements ICollections.
func (c *cache) AddHeadlineBanner(ctx context.Context, headline entity.Collection) error {
	return c.collection.AddHeadlineBanner(ctx, headline)
}

// SlicesOfCollections implements ICollections.
func (c *cache) SlicesOfCollections(ctx context.Context, position string, limit int) ([]entity.Collection, error) {
	return c.collection.SlicesOfCollections(ctx, position, limit)
}

// UploadCollectionImage implements ICollections.
func (c *cache) UploadCollectionImage(ctx context.Context, collectionID string, url string) error {
	return c.collection.UploadCollectionImage(ctx, collectionID, url)
}
