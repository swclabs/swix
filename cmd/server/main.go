// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"fmt"
	"log"

	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/internal/delivery"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := delivery.NewServer(addr)
	adapter := delivery.NewAdapter()

	if err := server.Connect(adapter); err != nil {
		log.Fatal(err)
	}
}
