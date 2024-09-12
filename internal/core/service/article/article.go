// Package article implements article
package article

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/repository/collections"
	"swclabs/swix/internal/core/repository/comments"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/infra/db"
)

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

// Article struct for article service
type Article struct {
	Blob        blob.IBlobStorage
	Collections collections.ICollections
	Comments    comments.ICommentRepository
}

// GetMessage implements IArticle.
func (p *Article) GetMessage(ctx context.Context, position string, limit int) (*dtos.Message, error) {
	collecs, err := p.Collections.GetMany(ctx, position, limit)
	if err != nil {
		return nil, err
	}
	var msg = dtos.Message{
		Position: position,
	}
	for _, collect := range collecs {
		var content dtos.HeadlineContent
		if err := json.Unmarshal([]byte(collect.Body), &content); err != nil {
			return nil, err
		}
		msg.Content = append(msg.Content, content.Content)
	}
	return &msg, nil
}

// UploadMessage implements IArticle.
func (p *Article) UploadMessage(ctx context.Context, message dtos.Message) error {
	for _, msg := range message.Content {
		json, _ := json.Marshal(dtos.HeadlineContent{
			Content: msg,
		})
		_, err := p.Collections.Create(ctx, entity.Collection{
			Position: message.Position,
			Body:     string(json),
		})
		if err != nil {
			return err
		}
	}
	return nil
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
	tx, err := db.NewTransaction(ctx)
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

		_, err := p.Comments.Insert(ctx, entity.Comments{
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
