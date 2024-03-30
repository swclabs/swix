package tools

import (
	"reflect"
	"runtime"
	"strings"
	"time"

	"swclabs/swipe-api/pkg/tools/logger"
)

type Job struct {
	function map[string]string
}

func NewJob() *Job {
	return &Job{
		function: make(map[string]string),
	}
}

func getName(input string) string {
	paths := strings.Split(input, "/")
	return paths[len(paths)-1]
}

func (job *Job) Scheduler(fn func(), _time time.Duration) {
	job.function[getName(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())] = _time.String()
	ticker := time.NewTicker(_time)
	go func() {
		for range ticker.C {
			fn()
		}
	}()
}

func (job *Job) Info() {
	logger.Banner("Launching a job scheduler with the following settings:")
	for fn, sche := range job.function {
		logger.CronLogger(fn, sche)
	}
}
