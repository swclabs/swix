package delivery

import (
	"time"

	"github.com/swclabs/swipe-api/internal/misc/cron"
	"github.com/swclabs/swipe-api/pkg/job"
)

func (account *_AccountManagementAdapter) _StartAccountManagementJob() {
	newJob := job.New()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}

func (adapter *_Adapter) _StartCommonJob() {
	newJob := job.New()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}
