package router

import "swclabs/swipecore/pkg/lib/worker"

type IRouter interface {
	Register(eng *worker.Engine)
}
