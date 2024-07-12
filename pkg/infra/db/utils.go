// Package db connect to database
package db

import (
	"errors"
	"swclabs/swipecore/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
)

func MigrateUp() error {
	const migrateURL = "file://pkg/migration/"
	databaseURL, err := utils.ConnectionURLBuilder("pg-migrate")
	if err != nil {
		return err
	}
	_migrate, err := migrate.New(migrateURL, databaseURL)
	if err != nil {
		return err
	}
	if err := _migrate.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func MigrateDown() error {
	const migrateURL = "file://pkg/db/migration/"
	databaseURL, err := utils.ConnectionURLBuilder("pg-migrate")
	if err != nil {
		return err
	}
	_migrate, err := migrate.New(migrateURL, databaseURL)
	if err != nil {
		return err
	}
	if err := _migrate.Down(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
