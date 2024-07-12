// Package router define tasks - queue
package router

import "swclabs/swipecore/pkg/lib/worker"

// IRouter interface for router objects
type IRouter interface {
	Register(eng *worker.Engine)
}
