// Package article implements article
package article

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/internal/core/repos/comments"
	"github.com/swclabs/swipex/internal/core/repos/news"
	"github.com/swclabs/swipex/internal/core/repos/products"
	"github.com/swclabs/swipex/internal/core/repos/stars"
	"github.com/swclabs/swipex/pkg/infra/blob"
	"github.com/swclabs/swipex/pkg/infra/db"
)

var _ = app.Service(New)

// New creates a new Article object
func New(
	blob blob.IBlobStorage,
	cmt comments.IComments,
	news news.INews,
) IArticle {
	return &Article{
		Blob:     blob,
		Comments: cmt,
		News:     news,
	}
}

// Article struct for article service
type Article struct {
	Blob     blob.IBlobStorage
	Comments comments.IComments
	News     news.INews
}

func (p *Article) UploadNews(ctx context.Context, newsDTO dtos.NewsDTO) (int64, error) {
	tx, err := db.NewTx(ctx)
	if err != nil {
		return -1, err
	}

	var newsRepo = news.New(tx)
	for _, card := range newsDTO.Cards {

		body, err := json.Marshal(card)
		if err != nil {
			return -1, err
		}

		_, err = newsRepo.Create(ctx, entity.News{
			Category: newsDTO.Category,
			Header:   newsDTO.Header,
			Body:     string(body),
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

func (p *Article) UploadNewsImage(ctx context.Context, newsID int64, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return err
	}

	resp, err := p.Blob.UploadImages(f)
	if err != nil {
		return err
	}

	return p.News.UploadNewsImage(ctx, newsID, resp.SecureURL)
}

func (p *Article) GetNews(ctx context.Context, category string, limit int) (*dtos.News, error) {
	newss, err := p.News.GetMany(ctx, category, limit)
	if err != nil {
		return nil, err
	}

	if len(newss) == 0 {
		return &dtos.News{}, nil
	}

	var nnews dtos.News
	nnews.Header = newss[0].Header

	for _, _collection := range newss {
		var body dtos.CardArticle
		if err := json.Unmarshal([]byte(_collection.Body), &body); err != nil {
			return nil, err
		}

		nnews.Cards = append(nnews.Cards, body)
	}
	return &nnews, nil
}

func (p *Article) GetComment(ctx context.Context, productID int64) ([]model.Comment, error) {
	return p.Comments.GetModelByProductID(ctx, productID)
}

// UploadComment implements IArticle.
func (p *Article) UploadComment(ctx context.Context, comment dtos.Comment) error {
	tx, err := db.NewTx(ctx)
	if err != nil {
		return err
	}

	var (
		stars    = stars.New(tx)
		products = products.New(tx)
		comments = comments.New(tx)
	)

	id, err := stars.Save(ctx, entity.Star{UserID: comment.UserID, ProductID: comment.ProductID})
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	if id == -1 {
		_ = tx.Rollback(ctx)
		return fmt.Errorf("you have already rated this product")
	}

	if err := products.Rating(ctx, comment.ProductID, comment.Star); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	_, err = comments.Insert(ctx, entity.Comment{
		StarID:      id,
		Content:     comment.Content,
		UserID:      comment.UserID,
		ProductID:   comment.ProductID,
		InventoryID: comment.InventoryID,
	})
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
