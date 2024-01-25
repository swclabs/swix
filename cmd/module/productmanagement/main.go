// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @kieranhoo
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

	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/internal/delivery"

	"github.com/urfave/cli/v2"

	_ "github.com/swclabs/swipe-api/docs"
)

var Command = []*cli.Command{
	{
		Name:    "worker",
		Aliases: []string{"w"},
		Usage:   "run worker handle tasks in queue",
		Action: func(_ *cli.Context) error {
			w := delivery.NewWorker()
			return w.Run(10)
		},
	},
	{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "run app server",
		Action: func(_ *cli.Context) error {
			addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
			server := delivery.NewServer(addr)
			adapter := delivery.NewProductManagementAdapter()

			return server.Connect(adapter)
		},
	},
}

func NewClient() *cli.App {
	newApp := &cli.App{
		Name:        "swipe",
		Usage:       "SwiftCart Project",
		Version:     "0.0.1",
		Description: "SwiftCart Product Management API server",
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
