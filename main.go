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
	"swclabs/swipecore/boot"
	"swclabs/swipecore/internal/core/service/common"
	"swclabs/swipecore/internal/webapi"
	"swclabs/swipecore/internal/webapi/controller"
	"swclabs/swipecore/internal/webapi/router"
)

// @title Swipe API documentation
// @version 1.0.0
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	var (
		commonService    = common.New()
		commonController = controller.NewCommon(commonService)
		commonRouter     = router.NewCommon(commonController)

		httpServer = webapi.NewServer(router.NewDocs(), commonRouter)
		adapt      = webapi.NewBaseAdapter(httpServer)
		server     = boot.NewServer()
	)

	log.Fatal(server.Connect(adapt))
}
