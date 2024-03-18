package delivery

import (
	"time"

	"github.com/swclabs/swipe-api/internal/helper/cron"
	"github.com/swclabs/swipe-api/internal/http"
	"github.com/swclabs/swipe-api/pkg/tools"
)

type IAdapter interface {
	Run(addr string) error
}

type _Adapter struct {
	server http.IServer
}

func NewAdapter() IAdapter {
	adapter := &_Adapter{
		server: http.New(),
	}
	// initialize cron job
	adapter._StartCommonJob()
	return adapter
}

// Run: run all services one server
func (adapter *_Adapter) Run(addr string) error {
	adapter.server.Bootstrap(
		http.AccountManagementModule,
		http.ProductManagementModule,
	)
	return adapter.server.Run(addr)
}

func (adapter *_Adapter) _StartCommonJob() {
	newJob := tools.NewJob()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}
