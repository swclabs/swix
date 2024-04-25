// Copyright 2023 Swipe. All rights reserved.
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

	"swclabs/swipe-api/boot"
	"swclabs/swipe-api/boot/adapter"
	_ "swclabs/swipe-api/docs"
	"swclabs/swipe-api/internal/config"
)

// @title Swipe API Documentation
// @version 1.0.0
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := boot.NewServer(addr)
	adapter := adapter.New(adapter.TypeBase)

	if err := server.Connect(adapter); err != nil {
		log.Fatal(err)
	}
}
