// Copyright 2023 Swipe. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author: - Ho Duc Hung : @kyeranyo
// 		   - Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

// THIS IS FILE USED TO CREATE SWAGGER DOCS ONLY
// PLEASE DO NOT EDIT, SEE cmd/main.go

package main

import (
	"log"
	"swclabs/swipecore/boot"
	"swclabs/swipecore/boot/adapter"
	_ "swclabs/swipecore/docs"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/core/service/common"
	"swclabs/swipecore/internal/http"
	"swclabs/swipecore/internal/http/controller"
	"swclabs/swipecore/internal/http/router"
	"swclabs/swipecore/pkg/lib/worker"
)

// @title Swipe API Documentation
// @version 1.0.0
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	var (
		env = config.LoadEnv()
		commonService = common.New(worker.NewClient(env))
		commonController = controller.NewCommon(commonService)
		commonRouter = router.NewCommon(commonController)
	
		httpServer = http.NewServer(commonRouter, router.NewDocs())
		adapt = adapter.NewBaseAdapter(httpServer)
		server = boot.NewServer(env)
	)

	log.Fatal(server.Connect(adapt))
}
