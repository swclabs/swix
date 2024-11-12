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
	"github.com/swclabs/swipex/internal/apis/container/healthcheck"
	"github.com/swclabs/swipex/internal/apis/server"
	service "github.com/swclabs/swipex/internal/core/service/healthcheck"

	_ "github.com/swclabs/swipex/docs"
)

// @title Swipe Public API v0.0.1
// @version 0.0.1
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	var (
		_service    = service.New()
		_controller = healthcheck.NewController(_service)
		_router     = healthcheck.NewRouter(_controller)

		mux    = server.NewServeMux()
		server = server.New(mux)
	)
	mux.Handle(_router)

	log.Fatal(server.Run())
}
