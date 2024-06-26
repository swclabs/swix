package db

import (
	"errors"
	"swclabs/swipecore/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
)

func MigrateUp() error {
	const migrateUrl = "file://pkg/db/migration/"
	databaseUrl, err := utils.ConnectionURLBuilder("pg-migrate")
	if err != nil {
		return err
	}
	_migrate, err := migrate.New(migrateUrl, databaseUrl)
	if err != nil {
		return err
	}
	if err := _migrate.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func MigrateDown() error {
	const migrateUrl = "file://pkg/db/migration/"
	databaseUrl, err := utils.ConnectionURLBuilder("pg-migrate")
	if err != nil {
		return err
	}
	_migrate, err := migrate.New(migrateUrl, databaseUrl)
	if err != nil {
		return err
	}
	if err := _migrate.Down(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
