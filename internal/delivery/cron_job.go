package delivery

import (
	"time"

	"github.com/swclabs/swipe-api/internal/helper/cron"
	"github.com/swclabs/swipe-api/pkg/tools"
)

func (account *_AccountManagementAdapter) _StartAccountManagementJob() {
	newJob := tools.NewJob()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}

func (adapter *_Adapter) _StartCommonJob() {
	newJob := tools.NewJob()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}
