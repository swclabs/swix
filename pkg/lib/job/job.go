// Package job represents background jobs
package job

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"swclabs/swipecore/pkg/lib/logger"
)

// Job struct
type Job struct {
	function map[string]string
}

// NewJob creates a new job
func NewJob() *Job {
	return &Job{
		function: make(map[string]string),
	}
}

func getName(input string) string {
	paths := strings.Split(input, "/")
	return paths[len(paths)-1]
}

// Scheduler schedules a job
func (job *Job) Scheduler(fn func(), _time time.Duration) {
	job.function[getName(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())] = _time.String()
	ticker := time.NewTicker(_time)
	go func() {
		for range ticker.C {
			fn()
		}
	}()
}

// Info prints the job scheduler settings
func (job *Job) Info() {
	logger.Info("launching a job scheduler with the following settings:")
	for fn, sche := range job.function {
		logger.Info(fmt.Sprintf("function: %s ==> sched: %s", fn, sche))
	}
}
