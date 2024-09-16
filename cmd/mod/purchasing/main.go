// Package main start purchasing module
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
		app := boot.NewApp(flags, boot.NewServer, apis.NewProductsAdapter)
		app.Run()
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
