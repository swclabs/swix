// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @ikierans
// - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

// THIS IS FILE USED TO CREATE SWAGGER DOCS ONLY
// PLEASE DO NOT EDIT, SEE cmd/main.go

package main

import (
	"fmt"
	"log"

	"swclabs/swiftcart/api"
	_ "swclabs/swiftcart/docs"
	"swclabs/swiftcart/internal/config"
)

// @title Swiftcart API Documentation
// @version 1.0.0
// @description This is a documentation for the Swiftcart API
// @host
// @basePath /
func main() {
	server := api.NewServer()
	if err := server.Run(fmt.Sprintf("%s:%s", config.Host, config.Port)); err != nil {
		log.Fatal(err)
	}
}
