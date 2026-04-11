package ui

import (
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	drawille "github.com/exrook/drawille-go"
)

// It scales the y-axis ratio 1 means perfect circles it gets wider with compression
// I suggest using around 0.4-0.7
var yCompression = 0.65

// Norman scale is 100x(100*(yCompression))

func drawCircle(canvas *drawille.Canvas, centerX, centerY, distance, scale float64, isSun bool) {
	thetaStep := 0.01
	radiusX := distance * scale
	radiusY := distance * yCompression * scale
	if isSun {
		radiusY = radiusX
	}
	for theta := 0.0; theta < math.Pi*2; theta += thetaStep {
		x := centerX + math.Cos(theta)*radiusX
		y := centerY + math.Sin(theta)*radiusY
		canvas.Set(int(x), int(y))
	}
}

func DrawOrbit(scale float64) string {
	canvas := drawille.NewCanvas()
	furthestPlanet := planets[len(planets)-1]
	width := int(furthestPlanet.orbitRadius * scale * 2)
	height := int(furthestPlanet.orbitRadius * yCompression * scale * 2)

	terminalPadding := 2
	terminalWidth := width/2 + terminalPadding*2
	terminalHeight := height/4 + terminalPadding*2
	centerX, centerY := float64(width)/2, float64(height)/2

	for _, p := range planets {
		drawCircle(&canvas, centerX, centerY, p.orbitRadius, scale, false)
	}

	sunRadius := 3.0
	for radius := sunRadius; radius > 0; radius -= 0.5 {
		drawCircle(&canvas, centerX, centerY, radius, scale, true)
	}

	sunBrailleRadius := sunRadius * scale
	sunTermMaxX := int(centerX+sunBrailleRadius) / 2
	sunTermMinX := int(centerX-sunBrailleRadius) / 2
	sunTermMinY := int(centerY-sunBrailleRadius) / 4
	sunTermMaxY := int(centerY+sunBrailleRadius) / 4

	buffer := make([][]string, terminalHeight)
	for y := 0; y < terminalHeight; y++ {
		buffer[y] = make([]string, terminalWidth)
		for x := 0; x < terminalWidth; x++ {
			buffer[y][x] = " "
		}
	}

	grayStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("0"))
	yellowStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("11"))

	lines := strings.Split(canvas.String(), "\n")
	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			bufY := y + terminalPadding
			bufX := x + terminalPadding
			if bufY >= terminalHeight || bufX >= terminalWidth {
				continue
			}

			if x >= sunTermMinX && x <= sunTermMaxX && y >= sunTermMinY && y <= sunTermMaxY {
				buffer[bufY][bufX] = yellowStyle.Render(char)
			} else {
				buffer[bufY][bufX] = grayStyle.Render(char)
			}
		}
	}

	for i, p := range planets {
		radiusX := p.orbitRadius * scale
		radiusY := p.orbitRadius * yCompression * scale

		// Exact planet coordinates on the orbit according to angle
		angle := planetAngles[i]
		brailleX := centerX + math.Cos(angle)*radiusX
		brailleY := centerY + math.Sin(angle)*radiusY
		termX := int(brailleX) / 2
		termY := int(brailleY) / 4

		style := lipgloss.NewStyle().Foreground(p.color)
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
