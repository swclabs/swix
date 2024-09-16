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
	"flag"
	"fmt"
	"swclabs/swix/boot"
	"swclabs/swix/internal/apis"
	"swclabs/swix/internal/workers"
	"swclabs/swix/pkg/lib/logger"
)

func main() {
	cmd := flag.String("start", "server", "start server or worker")
	flag.Usage = func() {
		fmt.Println("Usage: swipe [flags]")
		flag.PrintDefaults()
	}
	flag.Parse()

	switch *cmd {
	case "worker":
		flags := boot.Worker | boot.DebugMode
		app := boot.NewApp(flags, boot.NewWorker, workers.NewAdapter)
		app.Run()
	case "server":
		flags := boot.APIs | boot.DebugMode
		app := boot.NewApp(flags, boot.NewServer, apis.NewAdapter)
		app.Run()
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
