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
	"swclabs/swipecore/internal/http"
	"swclabs/swipecore/internal/workers"

	"swclabs/swipecore/boot"

	"github.com/urfave/cli/v2"
)

var Command = []*cli.Command{
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(_ *cli.Context) error {
			boot.PrepareFor(boot.WorkerConsume)
			app := boot.NewApp(boot.NewWorker, workers.NewAdapter)
			app.Run()
			return nil
		}, 
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run api server",
		Action: func(_ *cli.Context) error {
			app := boot.NewApp(boot.NewServer, http.NewAdapter)
			app.Run()
			return nil
		},
	},
}

func NewClient() *cli.App {
	newApp := &cli.App{
		Name:        "swipe",
		Usage:       "swipe",
		Version:     "0.0.1",
		Description: "Swipe API server cli",
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
