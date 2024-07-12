// Package types define internal types
package types

// IAdapter interface for adapter objects
type IAdapter interface {
	Run(addr string) error
	StartWorker(concurrency int) error
}
