package types

type IAdapter interface {
	Run(addr string) error
	StartWorker(concurrency int) error
}
