package domain

import "context"

type INewsletterRepository interface {
	Insert(ctx context.Context, newsletter Newsletter) error
	Get(ctx context.Context, limit int) ([]Newsletters, error)
	GetHomeBanner(ctx context.Context, limit int) ([]HomeBanners, error)
	InsertHomeBanner(ctx context.Context, homeBanner HomeBanners) error
}
