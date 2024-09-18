// Package router define tasks - queue
package router

import "swclabs/swix/pkg/lib/worker"

// IRouter interface for router objects
type IRouter interface {
	Register(eng worker.IEngine)
}
