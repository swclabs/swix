// Package db connect to database
package db

import "github.com/jackc/pgx/v5"

// CollectOneRow collects one row from the given Rows object
func CollectOneRow[T any](rows Rows) (T, error) {
	return pgx.CollectOneRow[T](rows, pgx.RowToStructByName[T])
}

// CollectRows collects all rows from the given Rows object
func CollectRows[T any](rows Rows) ([]T, error) {
	return pgx.CollectRows[T](rows, pgx.RowToStructByName[T])
}
