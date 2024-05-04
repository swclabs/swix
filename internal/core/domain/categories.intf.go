package domain

import "context"

type ICategoriesRepository interface {
	Insert(ctx context.Context, ctg *Categories) error
	GetLimit(ctx context.Context, limit string) ([]Categories, error)
}
