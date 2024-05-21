// Copyright 2023 Swipe. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author: - Ho Duc Hung : @kieranhoo
// 		   - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

// THIS IS FILE USED TO CREATE SWAGGER DOCS ONLY
// PLEASE DO NOT EDIT, SEE cmd/main.go

package main

import (
	"fmt"
	"log"

	"swclabs/swipecore/boot"
	"swclabs/swipecore/boot/adapter"
	_ "swclabs/swipecore/docs"
	"swclabs/swipecore/internal/config"
)

// @title Swipe API Documentation
// @version 1.0.0
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := boot.NewServer(addr)
	myAdapter := adapter.New(adapter.TypeBase)

	if err := server.Connect(myAdapter); err != nil {
		log.Fatal(err)
	}
}
