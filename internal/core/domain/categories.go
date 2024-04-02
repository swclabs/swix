package domain

import "context"

type CategoriesRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// Categories Table
type Categories struct {
	Id          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

type ICategoriesRepository interface {
	New(ctx context.Context, ctg *Categories) error
	GetAll(ctx context.Context) ([]Categories, error)
}
