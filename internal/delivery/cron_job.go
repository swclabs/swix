package delivery

import (
	"time"

	"github.com/swclabs/swipe-api/internal/misc/cron"
	"github.com/swclabs/swipe-api/pkg/job"
)

func _StartCommonJob() {
	newJob := job.New()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}
