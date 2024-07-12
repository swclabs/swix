// Package db connect to database
package db

import "github.com/jackc/pgx/v5"

func CollectOneRow[T any](rows Rows) (T, error) {
	return pgx.CollectOneRow[T](rows, pgx.RowToStructByName[T])
}

func CollectRows[T any](rows Rows) ([]T, error) {
	return pgx.CollectRows[T](rows, pgx.RowToStructByName[T])
}
