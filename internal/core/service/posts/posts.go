// Package posts implements posts
package posts

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/errors"
	"swclabs/swipecore/internal/core/repository/collections"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/utils"
)

// Posts struct for posts service
type Posts struct {
	Blob        blob.IBlobStorage
	Collections collections.ICollections
}

// New creates a new Posts object
func New(
	blob blob.IBlobStorage,
	collection collections.ICollections,
) IPostsService {
	return &Posts{
		Blob:        blob,
		Collections: collection,
	}
}

// SliceOfHeadlineBanner implements IPostsService.
func (p *Posts) SliceOfHeadlineBanner(
	ctx context.Context, position string, limit int) (*domain.HeadlineBannerSlices, error) {

	_collections, err := p.Collections.SlicesOfCollections(ctx, position, limit)
	if err != nil {
		return nil, errors.Service("get collections", err)
	}

	var headlineBanners domain.HeadlineBannerSlices
	headlineBanners.Position = position
	for _, collection := range _collections {
		var body domain.HeadlineBannerBody
		if err := json.Unmarshal([]byte(collection.Body), &body); err != nil {
			return nil, err
		}
		headlineBanners.Headlines = append(headlineBanners.Headlines,
			domain.HeadlineBannerSlicesBody{
				HeadlineBannerBody: body,
				ID:                 collection.ID,
				Created:            utils.HanoiTimezone(collection.Created),
			})
	}
	return &headlineBanners, nil
}

// UploadHeadlineBanner implements IPostsService.
func (p *Posts) UploadHeadlineBanner(ctx context.Context, banner domain.HeadlineBannerSchema) error {
	return p.Collections.AddHeadlineBanner(ctx, banner)
}

// SlicesOfCollections implements IPostsService.
func (p *Posts) SlicesOfCollections(
	ctx context.Context, position string, limit int) (*domain.CollectionSliceSchema, error) {
	collectionSlice, err := p.Collections.SlicesOfCollections(ctx, position, limit)
	if err != nil {
		return nil, err
	}

	var _collections domain.CollectionSliceSchema
	_collections.Position = collectionSlice[0].Position
	_collections.Headline = collectionSlice[0].Headline

	for _, _collection := range collectionSlice {
		var body domain.CollectionBody
		if err := json.Unmarshal([]byte(_collection.Body), &body); err != nil {
			return nil, err
		}
		_collections.CardBanner = append(_collections.CardBanner,
			domain.CollectionSliceBody{
				CollectionBody: body,
				ID:             _collection.ID,
				Created:        utils.HanoiTimezone(_collection.Created),
			})
	}
	return &_collections, nil
}

// UploadCollections implements IPostsService.
func (p *Posts) UploadCollections(
	ctx context.Context, banner domain.CollectionSchema) (int64, error) {
	return p.Collections.AddCollection(ctx, banner)
}

// UploadCollectionsImage implements IPostsService.
func (p *Posts) UploadCollectionsImage(
	ctx context.Context, cardBannerID string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := p.Blob.UploadImages(file)
	if err != nil {
		return err
	}
	return p.Collections.UploadCollectionImage(
		ctx, cardBannerID, resp.SecureURL)
}
