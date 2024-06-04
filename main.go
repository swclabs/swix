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
	"go.uber.org/fx"
	"swclabs/swipecore/boot"
	"swclabs/swipecore/boot/adapter"
	_ "swclabs/swipecore/docs"
)

// @title Swipe API Documentation
// @version 1.0.0
// @description This is a documentation for the Swipe API
// @host
// @basePath /
func main() {
	app := fx.New(
		boot.FxRestModule,
		fx.Provide(
			adapter.NewAdapter,
			boot.NewServer,
		),
		fx.Invoke(boot.StartServer),
	)
	app.Run()
}
