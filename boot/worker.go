// Package boot implement worker consume for swipe application
package boot

import (
	"strconv"
	"swclabs/swix/internal/config"
)

type _Worker struct {
	concurrency int
}

// NewWorker create new worker consume
func NewWorker() IServer {
	return &_Worker{
		concurrency: config.NumberOfWorker,
	}
}

func (w *_Worker) Bootstrap(core ICore) error {
	return core.Run(strconv.Itoa(w.concurrency))
}
