package model

import tea "charm.land/bubbletea/v2"

type model struct {
	Name string
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
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *model) View() tea.View {
	v := tea.NewView(m.Name)
	v.AltScreen = true
	return v
}
