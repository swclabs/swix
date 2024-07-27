// Package boot implement worker consume for swipe application
package boot

import (
	"strconv"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/types"
)

/*
 * Example:
 *	package main
 *	import (
 *		"log"
 *		"swclabs/swipecore/boot"
 *  )
 * 	func main() {
 *		w := boot.NewWorker()
 *		if err := w.Run(10); err != nil {
 *			log.Fatal(err)
 *		}
 *	}
 */

type _Worker struct {
	concurrency int
}

// NewWorker create new worker consume
func NewWorker() IServer {
	return &_Worker{
		concurrency: config.NumberOfWorker,
	}
}

func (w *_Worker) Connect(adapter types.IAdapter) error {
	return adapter.Run(strconv.Itoa(w.concurrency))
}
