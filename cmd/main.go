// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @ikierans
// - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

// RUN APPLICATION CLI, IF YOU DON'T WANT TO RUN CLI APP
// SEE: server/main.go and worker/main.go

package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"swclabs/swiftcart/api"

	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/urfave/cli/v2"

	_ "swclabs/swiftcart/docs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Command = []*cli.Command{
	{
		Name:    "migrate",
		Aliases: []string{"m"},
		Usage:   "migrate database",
		Action: func(c *cli.Context) error {
			const migrateUrl = "file://pkg/db/migration/"
			databaseUrl, err := utils.ConnectionURLBuilder("pg-migrate")
			if err != nil {
				return err
			}
			_migrate, err := migrate.New(migrateUrl, databaseUrl)
			if err != nil {
				return err
			}
			switch c.Args().First() {
			case "up":
				if err := _migrate.Up(); err != migrate.ErrNoChange {
					return err
				}
			case "down":
				if err := _migrate.Down(); err != migrate.ErrNoChange {
					return err
				}
			}
			return nil
		},
	},
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(_ *cli.Context) error {
			w := api.NewWorker(10)
			return w.Run()
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run api server",
		Action: func(_ *cli.Context) error {
			server := api.NewServer()
			return server.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
		},
	},
}

func NewClient() *cli.App {
	_app := &cli.App{
		Name:        "swiftcart",
		Usage:       "swiftcart",
		Version:     "0.0.1",
		Description: "swiftcart API server",
		Commands:    Command,
	}

	sort.Sort(cli.FlagsByName(_app.Flags))
	sort.Sort(cli.CommandsByName(_app.Commands))

	return _app
}

func main() {
	client := NewClient()

	if err := client.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
