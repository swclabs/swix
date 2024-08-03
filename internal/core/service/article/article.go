// Package article implements article
package article

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/repository/collections"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/infra/db"
	"swclabs/swipecore/pkg/lib/errors"
	"swclabs/swipecore/pkg/utils"
	"time"
)

// Article struct for article service
type Article struct {
	Blob        blob.IBlobStorage
	Collections collections.ICollections
}

// New creates a new Article object
func New(
	blob blob.IBlobStorage,
	collection collections.ICollections,
) IArticle {
	return &Article{
		Blob:        blob,
		Collections: collection,
	}
}

// SliceOfHeadlineBanner implements IArticle.
func (p *Article) SliceOfHeadlineBanner(
	ctx context.Context, position string, limit int) (*dtos.HeadlineBanners, error) {

	_collections, err := p.Collections.GetMany(ctx, position, limit)
	if err != nil {
		return nil, errors.Service("get collections", err)
	}

	var headlineBanners dtos.HeadlineBanners
	headlineBanners.Position = position
	for _, collection := range _collections {
		var body dtos.HeadlineBannerBody
		if err := json.Unmarshal([]byte(collection.Body), &body); err != nil {
			return nil, err
		}
		headlineBanners.Headlines = append(headlineBanners.Headlines,
			dtos.HeadlineBannerSlicesBody{
				HeadlineBannerBody: body,
				ID:                 collection.ID,
				Created:            utils.HanoiTimezone(collection.Created),
			})
	}
	return &headlineBanners, nil
}

// UploadHeadlineBanner implements IArticle.
func (p *Article) UploadHeadlineBanner(ctx context.Context, banner dtos.HeadlineBanner) error {
	body, err := json.Marshal(banner.Body)
	if err != nil {
		return err
	}

	return p.Collections.AddHeadlineBanner(ctx, entity.Collection{
		Position: banner.Position,
		Created:  time.Now().UTC(),
		Body:     string(body),
	})
}

// GetCarousels implements IArticle.
func (p *Article) GetCarousels(ctx context.Context, position string, limit int) (*dtos.Article, error) {
	collectionSlice, err := p.Collections.GetMany(ctx, position, limit)
	if err != nil {
		return nil, err
	}
	if len(collectionSlice) == 0 {
		return nil, fmt.Errorf("no collection found")
	}

	//var _collections dtos.Collections
	//_collections.Position = collectionSlice[0].Position
	//_collections.Headline = collectionSlice[0].Headline

	//for _, _collection := range collectionSlice {
	//	var body dtos.CollectionBody
	//	if err := json.Unmarshal([]byte(_collection.Body), &body); err != nil {
	//		return nil, err
	//	}
	//	_collections.CardBanner = append(_collections.CardBanner,
	//		dtos.CollectionSliceBody{
	//			CollectionBody: body,
	//			ID:             _collection.ID,
	//			Created:        utils.HanoiTimezone(_collection.Created),
	//		})
	//}

	var carousels dtos.Article
	carousels.Headline = collectionSlice[0].Headline

	for _, _collection := range collectionSlice {
		var body dtos.CardInArticle
		if err := json.Unmarshal([]byte(_collection.Body), &body); err != nil {
			return nil, err
		}
		carousels.Cards = append(carousels.Cards, body)
	}
	return &carousels, nil
}

// UploadArticle implements IArticle.
func (p *Article) UploadArticle(ctx context.Context, article dtos.UploadArticle) (int64, error) {
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return -1, err
	}
	var collection = collections.New(tx)
	for _, card := range article.Cards {
		body, err := json.Marshal(card)
		if err != nil {
			return -1, err
		}
		_, err = collection.Create(ctx, entity.Collection{
			Position: article.Position,
			Headline: article.Headline,
			Body:     string(body),
			Created:  time.Now().UTC(),
		})
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				return -1, err
			}
		}
	}
	if err = tx.Commit(ctx); err != nil {
		return -1, err
	}
	return 0, nil
}

// UploadCollectionsImage implements IArticle.
func (p *Article) UploadCollectionsImage(ctx context.Context, cardBannerID string, fileHeader *multipart.FileHeader) error {
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
