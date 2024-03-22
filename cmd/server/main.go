// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"fmt"
	"log"

	"github.com/swclabs/swipe-api/delivery"
	"github.com/swclabs/swipe-api/delivery/adapter"
	"github.com/swclabs/swipe-api/internal/config"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := delivery.NewServer(addr)
	adapter := adapter.NewAdapter()

	if err := server.Connect(adapter); err != nil {
		log.Fatal(err)
	}
}
