package main

import (
	"solcl/model"
	"solcl/ui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	ui.CalculatePlanetAngles()
	m := model.InitialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
