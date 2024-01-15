package delivery

import (
	"swclabs/swiftcart/internal/misc/cron"
	"swclabs/swiftcart/pkg/job"
	"time"
)

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

func (client *Client) scheduler() {
	newJob := job.New()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}

func (client *Client) ConnectTo(server IHttpServer) error {
	client.scheduler()
	return server.ListenOn(client.address)
}
