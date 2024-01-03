package http

type IHttpServer interface {
	ListenOn(addr string) error
}

type Client struct {
	address string
}

func NewClient(addr string) *Client {
	return &Client{
		address: addr,
	}
}

func (client *Client) ConnectTo(server IHttpServer) error {
	return server.ListenOn(client.address)
}
