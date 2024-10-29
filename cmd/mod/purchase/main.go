// Package main start purchase module
package main

import (
	"flag"
	"fmt"
	"log"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/container/purchase"
	purchaseWorker "swclabs/swipex/internal/workers/container/purchase"
	"swclabs/swipex/pkg/lib/logger"

	_ "swclabs/swipex/docs/purchase"
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
		log.Fatal(app.Run())
	case "server":
		app := app.Builder(purchase.New)
		log.Fatal(app.Run())
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
