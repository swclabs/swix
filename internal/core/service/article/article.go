// Package article implements article
package article

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/repos/comments"
	"swclabs/swipex/internal/core/repos/news"
	"swclabs/swipex/pkg/infra/blob"
	"swclabs/swipex/pkg/infra/db"
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

// GetComment implements IArticle.
//
//		{
//			"id": "11446498",
//			"level": "0", // level of comment, [0] is parent, [1] is child
//	     	"parent_id": "11446498", // parent id of comment
//			"created_at": "2012-12-12T10:53:43-08:00",
//			"username": "11446498",
//			"user_id": "11446498",
//			"product_id": "11446498",
//			"content": "@Aaron Levie these tigers are cool!",
//		  }
func (p *Article) GetComment(ctx context.Context, productID int64) ([]dtos.Comment, error) {
	comments, err := p.Comments.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}

	var comment = []dtos.Comment{}
	for _, cmt := range comments {

		var (
			level    int64
			parentID int64
		)

		if cmt.Level == 0 {
			// Add your code here
			level = cmt.Level
			parentID = 0
		}

		if cmt.Level == 1 {
			// Add your code here
			level = cmt.Level
			parentID = cmt.ParentID
		}
		
		comment = append(comment, dtos.Comment{
			ID:      cmt.ID,
			Level:   level,
			Content: []string{cmt.Content},
			// Username: cmt.Username,
			UserID:    cmt.UserID,
			ProductID: cmt.ProductID,
			ParentID:  parentID,
			Liked:     cmt.Liked,
			Disliked:  cmt.Disliked,
			// Add other fields here if needed
		})
	}

	return comment, nil
}

// UploadComment implements IArticle.
func (p *Article) UploadComment(ctx context.Context, comment dtos.Comment) error {
	for _, cmt := range comment.Content {

		_, err := p.Comments.Insert(ctx, entity.Comment{
			Level:     comment.Level,
			Content:   cmt,
			UserID:    comment.UserID,
			ProductID: comment.ProductID,
			ParentID:  comment.ParentID,
			Rating:    comment.Rating,
			Liked:     comment.Liked,
			Disliked:  comment.Disliked,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
