/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 *
 * * RUN APPLICATION CLI, IF YOU DON'T WANT TO RUN CLI APP
 * * SEE: server/main.go and worker/main.go
 */

package main

import (
	"log"
	"os"
	"sort"
	"swclabs/swipecore/boot"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/wapi"
	"swclabs/swipecore/internal/workers"

	"github.com/urfave/cli/v2"

	_ "swclabs/swipecore/boot/init"
	_ "swclabs/swipecore/docs"
)

var Command = []*cli.Command{
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(_ *cli.Context) error {
			if config.StageStatus == "prod" {
				boot.PrepareFor(boot.WorkerConsume | boot.ProdMode)
			} else {
				boot.PrepareFor(boot.WorkerConsume | boot.DebugMode)
			}
			app := boot.NewApp(boot.NewWorker, workers.NewAdapter)
			app.Run()
			return nil
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run app server",
		Action: func(_ *cli.Context) error {
			if config.StageStatus == "prod" {
				boot.PrepareFor(boot.RestAPI | boot.ProdMode)
			} else {
				boot.PrepareFor(boot.RestAPI | boot.DebugMode)
			}
			app := boot.NewApp(boot.NewServer, wapi.NewAccountManagementsAdapter)
			app.Run()
			return nil
		},
	},
}

func NewClient() *cli.App {
	newApp := &cli.App{
		Name:        "swipe",
		Usage:       "Swipe Project",
		Version:     "0.0.1",
		Description: "Swipe Account Management API server",
		Commands:    Command,
	}

	sort.Sort(cli.FlagsByName(newApp.Flags))
	sort.Sort(cli.CommandsByName(newApp.Commands))

	return newApp
}

func main() {
	client := NewClient()

	if err := client.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
