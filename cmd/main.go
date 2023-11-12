package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"example/komposervice/api"

	"example/komposervice/internal/config"
	"example/komposervice/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/urfave/cli/v2"

	_ "example/komposervice/docs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var Command = []*cli.Command{
	{
		Name:    "migrate",
		Aliases: []string{"m"},
		Usage:   "migrate database",
		Action: func(_ *cli.Context) error {
			const migrateUrl = "file://pkg/db/migration/"
			databaseUrl, err := utils.ConnectionURLBuilder("pg-migrate")
			if err != nil {
				return err
			}
			_migrate, err := migrate.New(migrateUrl, databaseUrl)
			if err != nil {
				return err
			}
			if err := _migrate.Up(); err != migrate.ErrNoChange {
				return err
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
		Name:        "komposervice",
		Usage:       "komposervice",
		Version:     "0.0.1",
		Description: "API server",
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
