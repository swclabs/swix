/*
Package boot implement api server for swipe application

Example:

	package main

	import (
		"fmt"
		"log"

		"swclabs/swipecore/boot"
		"swclabs/swipecore/boot/adapter"
		"swclabs/swipecore/internal/config"
	)

	func main() {
		addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
		server := boot.NewServer(addr)
		adapt := adapter.New(adapter.TypeBase)

		if err := server.Connect(adapt); err != nil {
			log.Fatal(err)
		}
	}
*/

package boot

import "swclabs/swipecore/boot/adapter"

type IServer interface {
	// Connect to adapter of other module
	Connect(adapter adapter.IAdapter) error
}

// struct server in project
//
// host:port - 127.0.0.1:8000
type _Server struct {
	address string //
}

// NewServer
//
// Example :host:port - 127.0.0.1:8000
func NewServer(addr string) IServer {
	return &_Server{
		address: addr,
	}
}

// Connect to module via adapter
//
// Example:
//
//	server := boot.NewServer("localhost:8000")
//	adapter := adapter.NewAdapter()
//	server.Connect(adapter)
func (server *_Server) Connect(adapter adapter.IAdapter) error {
	return adapter.Run(server.address)
}
