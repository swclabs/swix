// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @ikierans
// - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"log"

	"swclabs/swiftcart/api"
)

func main() {
	w := api.NewWorker(10)
	if err := w.Run(); err != nil {
		log.Fatal(err)
	}
}
