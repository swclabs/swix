package repository

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
	var newsletters []domain.Newsletters
	if err := news.conn.WithContext(ctx).
		Table(domain.NewsletterTable).
		Not("type=?", "home-banner").
		Find(&newsletters).
		Limit(limit).Error; err != nil {
		return nil, err
	}
	return newsletters, nil
}

// GetHomeBanner implements domain.INewsletterRepository.
func (news *Newsletter) GetHomeBanner(ctx context.Context, limit int) ([]domain.HomeBanners, error) {
	var newsletters []domain.Newsletters
	var homeBanners []domain.HomeBanners
	if err := news.conn.WithContext(ctx).
		Table(domain.NewsletterTable).
		Where("type=?", "home-banner").
		Find(&newsletters).
		Limit(limit).Error; err != nil {
		return nil, err
	}
	for _, newsletter := range newsletters {
		homeBanners = append(homeBanners, domain.HomeBanners{
			Name:     newsletter.Title,
			Subtitle: newsletter.SubTitle,
			Img:      newsletter.Image,
			Text:     newsletter.TextColor,
		})
	}
	return homeBanners, nil
}

// InsertHomeBanner implements domain.INewsletterRepository.
func (news *Newsletter) InsertHomeBanner(ctx context.Context, homeBanner domain.HomeBanners) error {
	return news.Insert(ctx, domain.Newsletter{
		Type:      "home-banner",
		Title:     homeBanner.Name,
		SubTitle:  homeBanner.Subtitle,
		TextColor: homeBanner.Text,
		Image:     homeBanner.Img,
	})
}
