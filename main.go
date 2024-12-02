package main

import (
	"fmt"
	"math"
	"time"
)

const width int = 42
const height int = 42

const halfWidth int = width / 2
const halfHeight int = height / 2

const dot byte = 46  // .
const dash byte = 45 // -
const star byte = 42 // *
const plus byte = 43 // +
const hash byte = 35 // #

func main() {
	screen := newScreen(dot)

	screen.drawDot(0, 0, hash)
	screen.drawDot(10, 0, hash)

	const scale float64 = 18.0

	for theta := 0.0; theta < 1000; theta += 0.1 {
		fmt.Println(halfWidth + int(math.Round(math.Cos(theta)*scale)))
		fmt.Println(halfHeight + int(math.Round(math.Sin(theta)*scale)))
		screen.clear(dot)

		screen.drawLine(
			21,
			21,
			halfWidth+int(math.Round(math.Cos(theta)*scale)),
			halfHeight+int(math.Round(math.Sin(theta)*scale)),
			star,
		)

		screen.drawDot(halfWidth, halfHeight, plus)
		screen.drawDot(
			halfWidth+int(math.Round(math.Cos(theta)*scale)),
			halfHeight+int(math.Round(math.Sin(theta)*scale)),
			plus,
		)
		renderWithBorders(screen)
		time.Sleep(time.Millisecond * 10)
	}

	renderWithBorders(screen)
}
