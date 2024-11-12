/**
 * package main
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
	"github.com/swclabs/swipex/internal/apis/container/products"
	"github.com/swclabs/swipex/pkg/lib/logger"

	_ "github.com/swclabs/swipex/docs/products"
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
	case "server":
		app := app.Builder(products.New)
		log.Fatal(app.Run())
	default:
		logger.Error("unknown flag: " + *cmd)
	}
}
