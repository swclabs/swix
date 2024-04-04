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
	data *domain.Newsletter
	conn *gorm.DB
}

func NewNewsletter() domain.INewsletterRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Newsletter{
		data: &domain.Newsletter{},
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
