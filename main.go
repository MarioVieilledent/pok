package main

import (
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

const black string = "\x1b[30m"
const red string = "\x1b[31m"
const green string = "\x1b[32m"
const yellow string = "\x1b[33m"
const blue string = "\x1b[34m"
const magenta string = "\x1b[35m"
const cyan string = "\x1b[36m"
const white string = "\x1b[37m"

const placeCursorTo0 string = "\x1b[0;0H"

func main() {
	setTitle("A game in go named pok")
	screen := newScreen(dot)

	screen.drawDot(0, 0, hash)
	screen.drawDot(10, 0, hash)

	const scale float64 = 18.0

	for theta := 0.0; theta < 1000; theta += 0.1 {
		screen.clear(dot)

		screen.drawLine(
			21,
			21,
			halfWidth+round(math.Cos(theta)*scale),
			halfHeight+round(math.Sin(theta)*scale),
			star,
		)

		screen.drawDot(halfWidth, halfHeight, plus)
		screen.drawDot(
			halfWidth+round(math.Cos(theta)*scale),
			halfHeight+round(math.Sin(theta)*scale),
			plus,
		)
		render(screen, green)
		time.Sleep(time.Millisecond * 10)
	}
}
