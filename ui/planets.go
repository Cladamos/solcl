package ui

import (
	"image/color"
	"math"
	"time"

	"charm.land/lipgloss/v2"
)

type planet struct {
	startAngle    float64
	realDistance  float64
	orbitalPeriod float64

	orbitRadius float64
	color       color.Color
}

// I think they are few packages do this math but I want to learn how its done :D

// Mean longitudes at J2000 epoch (degrees) — source: JPL Keplerian Elements
// https://ssd.jpl.nasa.gov/planets/approx_pos.html

var planets = []planet{
	{startAngle: 252.25, realDistance: 0.39, orbitalPeriod: 87.97, orbitRadius: 17.0, color: lipgloss.Color("241")},     // Mercury
	{startAngle: 181.98, realDistance: 0.72, orbitalPeriod: 224.70, orbitRadius: 28.0, color: lipgloss.Color("229")},    // Venus
	{startAngle: 100.47, realDistance: 1.00, orbitalPeriod: 365.26, orbitRadius: 38.0, color: lipgloss.Color("24")},     // Earth
	{startAngle: 355.43, realDistance: 1.52, orbitalPeriod: 686.98, orbitRadius: 53.0, color: lipgloss.Color("130")},    // Mars
	{startAngle: 34.33, realDistance: 5.20, orbitalPeriod: 4332.59, orbitRadius: 70.0, color: lipgloss.Color("137")},    // Jupiter
	{startAngle: 50.08, realDistance: 9.54, orbitalPeriod: 10759.22, orbitRadius: 78.0, color: lipgloss.Color("143")},   // Saturn
	{startAngle: 314.20, realDistance: 19.19, orbitalPeriod: 30688.50, orbitRadius: 86.0, color: lipgloss.Color("67")},  // Uranus
	{startAngle: 304.22, realDistance: 30.07, orbitalPeriod: 60182.00, orbitRadius: 100.0, color: lipgloss.Color("24")}, // Neptune
}

var planetAngles []float64

func CalculatePlanetAngles() {
	planetAngles = []float64{}
	referenceDate := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	daysSinceReference := time.Since(referenceDate).Hours() / 24

	for _, p := range planets {
		startRad := p.startAngle * math.Pi / 180
		angle := startRad + (2*math.Pi/p.orbitalPeriod)*daysSinceReference
		angle = math.Mod(angle, 2*math.Pi)
		if angle < 0 {
			angle += 2 * math.Pi
		}
		planetAngles = append(planetAngles, angle)
	}
}
