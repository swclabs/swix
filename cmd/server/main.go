// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"fmt"
	"log"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/internal/delivery"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := delivery.NewServer(addr)
	adapter := delivery.NewAdapter()

	if err := server.Connect(adapter); err != nil {
		log.Fatal(err)
	}
}
