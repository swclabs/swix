/*
Package boot implement worker consume for swipe application

Example:

	package main

	import (
		"log"

		"swclabs/swipecore/boot"
	)

	func main() {
		w := boot.NewWorker()
		if err := w.Run(10); err != nil {
			log.Fatal(err)
		}
	}
*/

package boot

import (
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/types"
)

type _Worker struct {
	concurrency int
}

// NewWorker create new worker consume
func NewWorker(env config.Env) IServer {
	return &_Worker{
		concurrency: 10,
	}
}

func (w *_Worker) Connect(adapter types.IAdapter) error {
	return adapter.StartWorker(w.concurrency)
}
