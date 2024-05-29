package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

func CacheSlicesOfCollections(
	collection ICollections, ctx context.Context, position string, limit int,
) ([]domain.Collection, error) {

	return collection.SlicesOfCollections(ctx, position, limit)
}
