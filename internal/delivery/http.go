package delivery

import (
	"swclabs/swiftcart/internal/http"
	"swclabs/swiftcart/internal/misc/cron"
	"swclabs/swiftcart/pkg/job"
	"time"
)

func InitCronJob() {
	newJob := job.New()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}

type IHttpServer interface {
	ListenOn(addr string) error
}

type Client struct {
	address string
}

func NewClient(addr string) *Client {
	InitCronJob()
	return &Client{
		address: addr,
	}
}

func Create(adapter ...func(*http.Server)) IHttpServer {
	var server = NewHttpServer()
	for _, adapt := range adapter {
		adapt(server.Server)
	}
	return server
}

func (client *Client) ConnectTo(server IHttpServer) error {
	return server.ListenOn(client.address)
}
