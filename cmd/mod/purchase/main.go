// Package main start purchase module
package main

import (
	"flag"
	"fmt"
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/container/purchase"
	purchaseWorker "swclabs/swix/internal/workers/container/purchase"
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
		app := app.Builder(purchaseWorker.New)
		_ = app.Run()
	case "server":
		app := app.Builder(purchase.New)
		_ = app.Run()
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
