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

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/apis"
	"github.com/swclabs/swipex/internal/cron"
	"github.com/swclabs/swipex/internal/workers"
	"github.com/swclabs/swipex/pkg/lib/logger"

	_ "github.com/swclabs/swipex/docs"
)

// @title Swipe Public API v0.0.1
// @version 0.0.1
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	cmd := flag.String("start", "server", "start server, worker or cronjob")
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
	case "cron":
		application := app.Builder(cron.NewApp)
		log.Fatal(application.Run())
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
