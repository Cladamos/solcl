package ui

import (
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	drawille "github.com/exrook/drawille-go"
)

// Real AU Distance: {0.39, 0.72, 1.0, 1.52, 5.2, 9.5, 19.2, 30.1}
var planetDistances = []float64{
	10.0,  //Mercury
	20.0,  //Venus
	30.0,  //Earth
	45.0,  //Mars
	70.0,  //Jupiter
	78.0,  //Saturn
	86.0,  //Uranus
	100.0, //Neptune
}

// It scales the y-axis ratio 1 means perfect circles it gets wider with compression
// I suggest using around 0.4-0.7
var yCompression = 0.55

// Norman scale is 100x(100*(yCompression))

func DrawOrbit(scale float64) string {
	canvas := drawille.NewCanvas()
	furthestPlanet := planetDistances[len(planetDistances)-1]
	width := int(furthestPlanet * scale * 2)
	height := int(furthestPlanet * yCompression * scale * 2)

	// All sides
	terminalPadding := 2
	terminalWidth := width/2 + terminalPadding*2
	terminalHeight := height/4 + terminalPadding*2
	centerX, centerY := float64(width)/2, float64(height)/2

	thetaStep := 0.005

	for _, p := range planetDistances {
		radiusX := p * scale
		radiusY := p * yCompression * scale
		// 1. Draw the static orbit path
		for theta := 0.0; theta < math.Pi*2; theta += thetaStep {
			x := centerX + math.Cos(theta)*radiusX
			y := centerY + math.Sin(theta)*radiusY
			canvas.Set(int(x), int(y))
		}
	}

	buffer := make([][]string, terminalHeight)
	for y := 0; y < terminalHeight; y++ {
		buffer[y] = make([]string, terminalWidth)
		for x := 0; x < terminalWidth; x++ {
			buffer[y][x] = " "
		}
	}

	grayStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	lines := strings.Split(canvas.String(), "\n")
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			buffer[y+terminalPadding][x+terminalPadding] = grayStyle.Render(char)
		}
	}

	for _, p := range planetDistances {
		radiusX := p * scale
		radiusY := p * yCompression * scale

		// Exact planet coordinates on the orbit according to angle
		brailleX := centerX + math.Cos(0)*radiusX
		brailleY := centerY + math.Sin(0)*radiusY
		termX := int(brailleX) / 2
		termY := int(brailleY) / 4

		style := lipgloss.NewStyle().Foreground(lipgloss.Color("208"))

		buffer[termY+terminalPadding][termX+terminalPadding] = style.Render("●")
	}

	var finalOutput strings.Builder
	for i, row := range buffer {
		for _, char := range row {
			finalOutput.WriteString(char)
		}
		if i < len(buffer)-1 {
			finalOutput.WriteString("\n")
		}
	}

	return finalOutput.String()
}
