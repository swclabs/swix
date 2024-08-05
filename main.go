// Copyright 2023 Swipe. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Ho Duc Hung : @kyeranyo
// Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

// THIS IS FILE USED TO CREATE SWAGGER DOCS ONLY
// PLEASE DO NOT EDIT, SEE cmd/main.go

package main

import (
	"log"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/service/base"
	"swclabs/swix/internal/webapi"
	"swclabs/swix/internal/webapi/controller"
	"swclabs/swix/internal/webapi/router"
)

// @title Swipe Public API v0.0.1
// @version 0.0.1
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	var (
		baseService    = base.New()
		baseController = controller.New(baseService)
		baseRouter     = router.New(baseController)

		httpServer = webapi.NewServer(baseRouter)
		adapt      = webapi.NewBaseAdapter(httpServer)
		server     = boot.NewServer()
	)

	log.Fatal(server.Connect(adapt))
}
