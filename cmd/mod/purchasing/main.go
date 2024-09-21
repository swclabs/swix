// Package main start purchasing module
package main

import (
	"flag"
	"fmt"
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/router"
	"swclabs/swix/internal/apis/server"
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
		// app := app.App(workers.NewWorkerNode)
		// _ = app.Run()
	case "server":
		app := app.App(func(purchase router.IPurchase) app.IApplication {
			mux := server.NewServeMux()
			mux.Handle(purchase)
			return server.New(mux)
		})
		_ = app.Run()
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
