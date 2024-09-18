// Package main start purchasing module
package main

import (
	"flag"
	"fmt"
	"swclabs/swix/app"
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
		app := app.App(workers.NewWorkerNode)
		_ = app.Run()
	case "server":
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
