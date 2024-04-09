package repo

import (
	"context"
	"log"
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/pkg/db"
	"swclabs/swipe-api/pkg/db/queries"

	"gorm.io/gorm"
)

type Newsletter struct {
	conn *gorm.DB
}

func NewNewsletter() domain.INewsletterRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Newsletter{
		conn: _conn,
	}
}

func (news *Newsletter) Insert(ctx context.Context, newsletter domain.Newsletter) error {
	return db.SafeWriteQuery(
		ctx,
		news.conn,
		queries.InsertIntoNewsletter,
		newsletter.Type, newsletter.Title, newsletter.SubTitle, newsletter.Description, newsletter.Image, newsletter.TextColor,
	)
}

func (news *Newsletter) Get(ctx context.Context, limit int) ([]domain.Newsletters, error) {
	var newsletter []domain.Newsletters
	if err := news.conn.Table(domain.NewsletterTable).Find(&newsletter).Limit(limit).Error; err != nil {
		return nil, err
	}
	return newsletter, nil
}
