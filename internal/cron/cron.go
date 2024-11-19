package cron

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/cron/register"
	"github.com/swclabs/swipex/internal/cron/server"
)

func NewApp() app.IApplication {
	cron := server.New()
	register.Statistic(cron)
	return cron
}
