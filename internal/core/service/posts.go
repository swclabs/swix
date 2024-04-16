package service

import (
	"context"
	"mime/multipart"
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/repo"
	"swclabs/swipe-api/pkg/cloud"
)

type Posts struct {
	category   domain.ICategoriesRepository
	product    domain.IProductRepository
	newsletter domain.INewsletterRepository
}

func NewPost() domain.IPostsService {
	return &Posts{
		category:   repo.NewCategories(),
		product:    repo.NewProducts(),
		newsletter: repo.NewNewsletter(),
	}
}

// GetHomeBanner implements domain.IPostsService.
func (p *Posts) GetHomeBanner(ctx context.Context) ([]domain.HomeBanners, error) {
	panic("unimplemented")
}

// GetNewsletter implements domain.IPostsService.
func (p *Posts) GetNewsletter(ctx context.Context, limit int) ([]domain.Newsletters, error) {
	return p.newsletter.Get(ctx, limit)
}

// UploadHomeBanner implements domain.IPostsService.
func (p *Posts) UploadHomeBanner(ctx context.Context, data *domain.HomeBanners) error {
	panic("unimplemented")
}

// UploadNewsletter implements domain.IPostsService.
func (p *Posts) UploadNewsletter(ctx context.Context, news domain.Newsletter, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	resp, err := cloud.UploadImages(cloud.Connection(), file)
	if err != nil {
		return err
	}
	news.Image = resp.SecureURL
	return repo.NewNewsletter().Insert(context.Background(), news)
}
