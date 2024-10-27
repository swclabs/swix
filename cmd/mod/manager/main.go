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
	"swclabs/swipex/internal/apis/container/authentication"
	authWorker "swclabs/swipex/internal/workers/container/authentication"
	"swclabs/swipex/pkg/lib/logger"

	_ "swclabs/swipex/app/init"
	_ "swclabs/swipex/docs"
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
		app := app.Builder(authWorker.New)
		log.Fatal(app.Run())
	case "server":
		app := app.Builder(authentication.New)
		log.Fatal(app.Run())
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
