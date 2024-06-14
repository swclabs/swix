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
	"swclabs/swipecore/boot/adapter"
	_ "swclabs/swipecore/docs"

	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var Command = []*cli.Command{
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(_ *cli.Context) error {
			app := fx.New(
				boot.FxWorkerModule,
				fx.Provide(
					boot.NewWorker,
				),
				fx.Invoke(boot.StartWorker),
			)
			app.Run()
			return nil
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run app server",
		Action: func(_ *cli.Context) error {
			app := fx.New(
				boot.FxRestModule,
				fx.Provide(
					adapter.NewAccountManagements,
					boot.NewServer,
				),
				fx.Invoke(boot.StartServer),
			)
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
