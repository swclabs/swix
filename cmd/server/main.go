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
	"swclabs/swiftcart/internal/app"
	"swclabs/swiftcart/internal/config"
)

func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	client := app.NewClient(addr)
	ginFrameworkAdapter := app.NewGinAdapter()

	if err := client.ConnectTo(ginFrameworkAdapter); err != nil {
		log.Fatal(err)
	}
}
