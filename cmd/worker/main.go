// Author:
// - Ho Duc Hung : @kieranhoo
// - Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"log"

	"github.com/swclabs/swipe-api/internal/delivery"
)

func main() {
	w := delivery.NewWorker()
	if err := w.Run(10); err != nil {
		log.Fatal(err)
	}
}
