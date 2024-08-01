// Package main contains the entry point of the base server.
package main

import (
	"log"
	"swclabs/swipecore/boot"
	"swclabs/swipecore/internal/mod/base"
)

func main() {
	var (
		server     = boot.NewServerWithAddress(":5001")
		grpcEngine = base.NewBaseServer()
		adapter    = base.New(grpcEngine)
	)
	log.Fatal(server.Connect(adapter))
}
