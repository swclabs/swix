package posts

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/collections"
	"swclabs/swipecore/pkg/blob"
)

type Posts struct {
	Collections collections.ICollections
}

func New(collection *collections.Collections) *Posts {
	return &Posts{
		Collections: collection,
	}
}

// SliceOfHeadlineBanner implements domain.IPostsService.
func (p *Posts) SliceOfHeadlineBanner(
	ctx context.Context, position string, limit int) (*domain.HeadlineBannerSlice, error) {

	collections, err := p.Collections.SlicesOfCollections(ctx, position, limit)
	if err != nil {
		return nil, err
	}

	var headlineBanners domain.HeadlineBannerSlice
	headlineBanners.Position = collections[0].Position
	for _, collection := range collections {
		var body domain.HeadlineBannerBody
		if err := json.Unmarshal([]byte(collection.Body), &body); err != nil {
			return nil, err
		}
		headlineBanners.Headlines = append(headlineBanners.Headlines, domain.HeadlineBannerSliceBody{
			HeadlineBannerBody: body,
			Id:                 collection.Id,
			Created:            collection.Created,
		})
	}
	return &headlineBanners, nil
}

// UploadHeadlineBanner implements domain.IPostsService.
func (p *Posts) UploadHeadlineBanner(ctx context.Context, banner domain.HeadlineBannerType) error {
	return p.Collections.AddHeadlineBanner(ctx, banner)
}

// SlicesOfCollections implements domain.IPostsService.
func (p *Posts) SlicesOfCollections(ctx context.Context, position string, limit int) (*domain.Collections, error) {
	collections, err := p.Collections.SlicesOfCollections(ctx, position, limit)
	if err != nil {
		return nil, err
	}

	var _collections domain.Collections
	_collections.Position = collections[0].Position
	_collections.Headline = collections[0].Headline

	for _, _collection := range collections {
		var body domain.CollectionBody
		if err := json.Unmarshal([]byte(_collection.Body), &body); err != nil {
			return nil, err
		}
		_collections.CardBanner = append(_collections.CardBanner, domain.CollectionsBody{
			CollectionBody: body,
			Id:             _collection.Id,
			Created:        _collection.Created,
		})
	}
	return &_collections, nil
}

// UploadCollections implements domain.IPostsService.
func (p *Posts) UploadCollections(ctx context.Context, banner domain.CollectionType) (int64, error) {
	return p.Collections.AddCollection(ctx, banner)
}

// UploadCollectionsImage implements domain.IPostsService.
func (p *Posts) UploadCollectionsImage(
	ctx context.Context, cardBannerId string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := blob.UploadImages(blob.Connection(), file)
	if err != nil {
		return err
	}
	return p.Collections.UploadCollectionImage(ctx, cardBannerId, resp.SecureURL)
}
