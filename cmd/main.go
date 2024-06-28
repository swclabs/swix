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

	"github.com/urfave/cli/v2"
	"swclabs/swipecore/boot"

	_ "swclabs/swipecore/boot/init"
)

var Command = []*cli.Command{
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(_ *cli.Context) error {
			app := boot.NewWorkerApp()
			app.Run()
			return nil
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run api server",
		Action: func(_ *cli.Context) error {
			app := boot.NewRestApp(http.NewAdapter)
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
