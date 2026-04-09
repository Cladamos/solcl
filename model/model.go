package model

import (
	"solcl/ui"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type model struct {
	width  int
	height int
	Name   string
}

func InitialModel() *model {
	return &model{
		Name: "Initial Model",
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *model) View() tea.View {
	orbit := ui.DrawOrbit()
	centeredOrbit := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, orbit)
	v := tea.NewView(centeredOrbit)
	v.AltScreen = true
	return v
}
