package repository

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type NewsletterMock struct {
	mock.Mock
}

var _ domain.INewsletterRepository = (*NewsletterMock)(nil)

// Get implements domain.INewsletterRepository.
func (n *NewsletterMock) Get(ctx context.Context, limit int) ([]domain.Newsletters, error) {
	args := n.Called(ctx, limit)
	return args.Get(0).([]domain.Newsletters), args.Error(1)
}

// GetHomeBanner implements domain.INewsletterRepository.
func (n *NewsletterMock) GetHomeBanner(ctx context.Context, limit int) ([]domain.HomeBanners, error) {
	args := n.Called(ctx, limit)
	return args.Get(0).([]domain.HomeBanners), args.Error(1)
}

// Insert implements domain.INewsletterRepository.
func (n *NewsletterMock) Insert(ctx context.Context, newsletter domain.Newsletter) error {
	args := n.Called(ctx, newsletter)
	return args.Error(0)
}

// InsertHomeBanner implements domain.INewsletterRepository.
func (n *NewsletterMock) InsertHomeBanner(ctx context.Context, homeBanner domain.HomeBanners) error {
	args := n.Called(ctx, homeBanner)
	return args.Error(0)
}
