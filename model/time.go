package model

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type tickMsg struct {
}

type speed struct {
	name  string
	t     time.Duration
	index int
}

var speeds = []speed{
	{"1s / s", time.Second, 0},
	{"1m / s", time.Minute, 1},
	{"1h / s", time.Hour, 2},
	{"1d / s", time.Hour * 24, 3},
	{"1w / s", time.Hour * 24 * 7, 4},
	{"1mo / s", time.Hour * 24 * 30, 5},
	{"1y / s", time.Hour * 24 * 365, 6},
}

func tickTime() tea.Cmd {
	return tea.Tick(time.Second, func(_ time.Time) tea.Msg {
		return tickMsg{}
	})
}
