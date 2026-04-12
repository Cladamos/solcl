package main

import (
	"github.com/cladamos/solcl/model"

	tea "charm.land/bubbletea/v2"
)

func main() {
	m := model.InitialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
