package ui

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Real AU Distance: {0.39, 0.72, 1.0, 1.52, 5.2, 9.5, 19.2, 30.1}
type planet struct {
	distance float64
	color    color.Color
	hasRing  bool
}

var planets = []planet{
	{distance: 17.0, color: lipgloss.Color("241"), hasRing: false}, // Mercury
	{distance: 28.0, color: lipgloss.Color("229"), hasRing: false}, // Venus
	{distance: 38.0, color: lipgloss.Color("24"), hasRing: false},  // Earth
	{distance: 53.0, color: lipgloss.Color("130"), hasRing: false}, // Mars
	{distance: 70.0, color: lipgloss.Color("137"), hasRing: false}, // Jupiter
	{distance: 78.0, color: lipgloss.Color("143"), hasRing: true},  // Saturn
	{distance: 86.0, color: lipgloss.Color("67"), hasRing: false},  // Uranus
	{distance: 100.0, color: lipgloss.Color("24"), hasRing: false}, // Neptune
}
