package delivery

import (
	"swclabs/swiftcart/internal/misc/cron"
	"swclabs/swiftcart/pkg/job"
	"time"
)

func _StartCommonJob() {
	newJob := job.New()
	go newJob.Scheduler(cron.Ping, 5*time.Second)

	newJob.Info()
}
