// account management adapter, base on base.adapt.go

package delivery

import (
	"time"

	"github.com/swclabs/swipe-api/internal/helper/cron"
	"github.com/swclabs/swipe-api/internal/http"
	"github.com/swclabs/swipe-api/pkg/tools"
)

type _AccountManagementAdapter struct {
	server http.IServer
}

func NewAccountManagementAdapter() IAdapter {
	account := &_AccountManagementAdapter{
		server: http.New(),
	}
	// initialize cron job
	account._StartAccountManagementJob()
	return account
}

func (account *_AccountManagementAdapter) Run(addr string) error {
	account.server.Bootstrap(http.AccountManagementModule)
	return account.server.Run(addr)
}

func (account *_AccountManagementAdapter) _StartAccountManagementJob() {
	newJob := tools.NewJob()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}
