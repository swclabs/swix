// Copyright 2023 Swiftcart. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// Description: This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"log"

	"github.com/swclabs/swipe-server/internal/delivery"
)

func main() {
	w := delivery.NewWorker()
	if err := w.Run(10); err != nil {
		log.Fatal(err)
	}
}
