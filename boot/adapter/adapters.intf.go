package adapter

// IAdapter interface, used to connect to server instance
type IAdapter interface {
	Run(addr string) error
}
