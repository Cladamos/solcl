package ui

import (
	"math"

	drawille "github.com/exrook/drawille-go"
)

func DrawOrbit(width, height int) string {
	canvas := drawille.NewCanvas()

	brailleScale := 1.5
	centerX, centerY := 80.0*brailleScale, 40.0*brailleScale
	radiusX, radiusY := 0.0, 0.0

	thetaStep := 0.005

	// Not close to accurate but looks fine I guess
	planetDistances := []float64{
		6.0,  // Mercury
		10.0, // Venus
		14.0, // Earth
		21.0, // Mars
		48.0, // Jupiter
		58.0, // Saturn
		70.0, // Uranus
		80.0, // Neptune
	}

	for _, planetDistance := range planetDistances {
		radiusX = planetDistance * brailleScale
		radiusY = planetDistance / 2 * brailleScale
		// 1. Draw the static orbit path
		for theta := 0.0; theta < math.Pi*2; theta += thetaStep {
			x := centerX + math.Cos(theta)*radiusX
			y := centerY + math.Sin(theta)*radiusY
			canvas.Set(int(x), int(y))
		}
	}
	return canvas.String()
}
