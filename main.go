// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

// THIS IS FILE USED TO CREATE SWAGGER DOCS ONLY
// PLEASE DO NOT EDIT, SEE cmd/main.go

package main

import (
	"fmt"
	"log"
	_ "github.com/swclabs/swipe-api/docs"
	"github.com/swclabs/swipe-api/internal/delivery"
	"github.com/swclabs/swipe-api/internal/config"
)

// @title Swiftcart API Documentation
// @version 1.0.0
// @description This is a documentation for the Swiftcart API
// @host
// @basePath /
func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := delivery.NewServer(addr)
	adapter := delivery.NewAdapter()

	if err := server.Connect(adapter); err != nil {
		log.Fatal(err)
	}
}
