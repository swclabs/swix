package server

import (
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/swclabs/swipex/internal/config"
)

type ICron interface {
	Run() error
	Register(cronspec string, task *asynq.Task, opts ...asynq.Option)
}

func New() ICron {
	return &cron{
		scheduler: asynq.NewScheduler(
			asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)},
			&asynq.SchedulerOpts{Location: time.Local},
		),
	}
}

type cron struct {
	scheduler *asynq.Scheduler
}

// Register implements ICron.
func (c *cron) Register(cronspec string, task *asynq.Task, opts ...asynq.Option) {
	_, err := c.scheduler.Register(cronspec, task, opts...)
	if err != nil {
		log.Fatal(err)
	}
}

// Run implements ICron.
func (c *cron) Run() error {
	return c.scheduler.Run()
}
