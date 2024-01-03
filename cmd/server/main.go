// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @ikierans
// - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"fmt"
	"log"
	"swclabs/swiftcart/delivery/http"
	"swclabs/swiftcart/internal/config"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := http.NewClient(addr)
	ginFrameworkAdapter := http.NewGinAdapter()

	if err := client.ConnectTo(ginFrameworkAdapter); err != nil {
		log.Fatal(err)
	}
}
