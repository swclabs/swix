package adapter

// IAdapter interface, used to connect to server instance
type IAdapter interface {
	Run(addr string) error
}

func New(types string) IAdapter {
	switch types {
	case TypeBase:
		return _NewAdapter()
	case TypeAccountManagement:
		return _NewAccountManagement()
	case TypeProducts:
		return _NewProducts()
	}
	return _NewAdapter()
}
