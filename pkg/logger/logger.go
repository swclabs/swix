package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func Banner(msg string) {
	styles := NewInfoStyleLogger()

	var gwlogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "[SWIPE]",
	})
	gwlogger.SetStyles(styles)
	gwlogger.Info(msg)
}

func CronLogger(function, schedule string) {
	styles := NewInfoStyleLogger()

	styles.Keys["function"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["function"] = lipgloss.NewStyle().Bold(true)
	styles.Keys["schedule"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["schedule"] = lipgloss.NewStyle().Bold(true)

	var gwlogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "[SWIPE-cron]",
	})
	gwlogger.SetStyles(styles)
	gwlogger.Info("-", "function", function, "schedule", schedule)
}

func Queue(queue string, priority int) {
	styles := NewInfoStyleLogger()

	styles.Keys["name"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["name"] = lipgloss.NewStyle().Bold(true)
	styles.Keys["priority"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["priority"] = lipgloss.NewStyle().Bold(true)

	var gwlogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "[SWIPE-queue]",
	})
	gwlogger.SetStyles(styles)
	gwlogger.Info("-", "name", queue, "priority", priority)
}

func Broker(name, host string) {
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: "86",
			Dark:  "86",
		}).
		Foreground(lipgloss.Color("0"))

	styles.Keys[name] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values[name] = lipgloss.NewStyle().Bold(true)

	var gwlogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "[SWIPE-broker]",
	})
	gwlogger.SetStyles(styles)
	gwlogger.Info("-", name, host)
}

func HandleFunc(typename string, handle string) {
	styles := NewInfoStyleLogger()

	styles.Keys["typename"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["typename"] = lipgloss.NewStyle().Bold(true)
	styles.Keys["handler"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["handler"] = lipgloss.NewStyle().Bold(true)

	var gwlogger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
		Prefix:          "[SWIPE-worker]",
	})
	gwlogger.SetStyles(styles)
	gwlogger.Info("-", "typename", typename, "handler", handle)
}
