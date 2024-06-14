// Package categories
// Author: Duc Hung Ho @kyeranyo
package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Categories struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Categories {
	return &Categories{conn: conn}
}

// Insert implements domain.ICategoriesRepository.
func (category *Categories) Insert(ctx context.Context, ctg domain.Categories) error {
	return db.SafePgxWriteQuery(
		ctx, category.conn, InsertIntoCategory, ctg.Name, ctg.Description)
}

// GetLimit implements domain.ICategoriesRepository.
func (category *Categories) GetLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	rows, err := category.conn.Query(ctx, SelectCategoryLimit, limit)
	if err != nil {
		return nil, err
	}
	categories, err := pgx.CollectRows[domain.Categories](rows, pgx.RowToStructByName[domain.Categories])
	if err != nil {
		return nil, err
	}
	return categories, nil
}
