package model

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type tickMsg struct {
}

type moveTimeMsg struct {
	time time.Time
}

func tickEveryHour() tea.Cmd {
	return tea.Tick(time.Hour, func(_ time.Time) tea.Msg {
		return tickMsg{}
	})
}

func moveTime(currentTime time.Time, amount time.Duration) tea.Msg {
	newTime := currentTime.Add(amount)
	return moveTimeMsg{time: newTime}
}
