// Package main contains the entry point of the base server.
package main

import (
	"log"
	"swclabs/swix/boot"
	"swclabs/swix/internal/mod/base"
)

func main() {
	var (
		server     = boot.NewServerWithAddress(":5001")
		grpcEngine = base.NewBaseServer()
		adapter    = base.New(grpcEngine)
	)
	log.Fatal(server.Connect(adapter))
}
