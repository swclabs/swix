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
	"log"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis"
	"swclabs/swipex/internal/workers"
	"swclabs/swipex/pkg/lib/logger"

	_ "swclabs/swipex/docs"
)

// @title Swipe Public API v0.0.1
// @version 0.0.1
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	cmd := flag.String("start", "server", "start server or worker")
	flag.Usage = func() {
		fmt.Println("Usage: swipe [flags]")
		flag.PrintDefaults()
	}
	flag.Parse()

	switch *cmd {
	case "worker":
		application := app.Builder(workers.NewApp)
		log.Fatal(application.Run())
	case "server":
		application := app.Builder(apis.NewApp)
		log.Fatal(application.Run())
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
