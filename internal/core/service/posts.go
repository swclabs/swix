package service

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository"
	"swclabs/swipecore/pkg/blob"
)

type Posts struct {
	CardBanner domain.ICollections
}

func NewPost() domain.IPostsService {
	return &Posts{
		CardBanner: repository.NewCardBannerCollection(),
	}
}

// SlicesOfCollections implements domain.IPostsService.
func (p *Posts) SlicesOfCollections(ctx context.Context, position string, limit int) (*domain.Collections, error) {
	return p.CardBanner.SlicesOfCollections(ctx, position, limit)
}

// UploadCollections implements domain.IPostsService.
func (p *Posts) UploadCollections(ctx context.Context, banner domain.CollectionType) (int64, error) {
	return p.CardBanner.AddCollection(ctx, banner)
}

// UploadCollectionsImage implements domain.IPostsService.
func (p *Posts) UploadCollectionsImage(ctx context.Context, cardBannerId string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := blob.UploadImages(blob.Connection(), file)
	if err != nil {
		return err
	}
	return p.CardBanner.UploadCollectionImage(ctx, cardBannerId, resp.SecureURL)
}
