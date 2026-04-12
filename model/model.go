package model

import (
	"solcl/ui"
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	width  int
	height int
	scale  float64
	time   time.Time

	isHelpHidden bool
	keys         keyMap
	help         help.Model
}
type keyMap struct {
	Quit         key.Binding
	Plus         key.Binding
	Minus        key.Binding
	Hide         key.Binding
	Reset        key.Binding
	TimeForward  key.Binding
	TimeBackward key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Plus, k.Minus, k.Hide, k.Reset}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Plus, k.Minus},
		{k.Hide, k.Reset},
	}
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
	),
	Plus: key.NewBinding(
		key.WithKeys("+", "up"),
		key.WithHelp("↑/+", "scale up"),
	),
	Minus: key.NewBinding(
		key.WithKeys("-", "down"),
		key.WithHelp("↓/-", "scale down"),
	),
	Hide: key.NewBinding(
		key.WithKeys("h", "H"),
		key.WithHelp("h", "hide help"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r", "R"),
		key.WithHelp("r", "reset time"),
	),
	TimeForward: key.NewBinding(
		key.WithKeys("right"),
	),
	TimeBackward: key.NewBinding(
		key.WithKeys("left"),
	),
}

func InitialModel() *model {
	currentTime := time.Now()
	ui.CalculatePlanetAngles(currentTime)
	return &model{
		keys:         keys,
		help:         help.New(),
		isHelpHidden: false,
		scale:        0.8,
		time:         currentTime,
	}
}

func (m *model) Init() tea.Cmd {
	return tickEveryHour()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.time = m.time.Add(time.Hour)
		ui.CalculatePlanetAngles(m.time)
		return m, tickEveryHour()
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.SetWidth(msg.Width)
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Plus):
			if m.scale < 1.5 {
				m.scale += 0.1
			}
		case key.Matches(msg, m.keys.Minus):
			if m.scale > 0.6 {
				m.scale -= 0.1
			}
		case key.Matches(msg, m.keys.Hide):
			m.isHelpHidden = !m.isHelpHidden
		case key.Matches(msg, m.keys.Reset):
			m.time = time.Now()
			ui.CalculatePlanetAngles(m.time)
		case key.Matches(msg, m.keys.TimeForward):
			m.time = m.time.Add(time.Hour * 24 * 1)
			ui.CalculatePlanetAngles(m.time)
		case key.Matches(msg, m.keys.TimeBackward):
			m.time = m.time.Add(-time.Hour * 24 * 1)
			ui.CalculatePlanetAngles(m.time)
		}
	}
	return m, nil
}

func (m *model) View() tea.View {
	orbits := ui.DrawOrbit(m.scale)
	timeStr := "↞ " + m.time.Format("2006-01-02 15:04:05") + " ↠\n"
	joined := lipgloss.JoinVertical(lipgloss.Center, orbits, timeStr)

	if !m.isHelpHidden {
		info := lipgloss.NewStyle().Foreground(lipgloss.Color("#4A4A4A")).Render("Ecliptic North (top-down) - Vernal equinox (right)")
		helpView := m.help.View(m.keys)
		joined = lipgloss.JoinVertical(lipgloss.Center, joined, info, helpView)
	}
	app := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, joined)

	v := tea.NewView(app)
	v.AltScreen = true
	return v
}
