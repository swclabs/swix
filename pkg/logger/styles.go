package logger

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func NewInfoStyleLogger() *log.Styles {
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: "86",
			Dark:  "86",
		}).
		Foreground(lipgloss.Color("0"))
	return styles
}
