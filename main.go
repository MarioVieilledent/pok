package main

import (
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

const speed = 0.05

func main() {
	setTitle("pok")
	screen := newScreen(dot)
	camera := newCamera()
	_ = screen
	_ = camera

	inputChannel := make(chan string)
	go listenInput(inputChannel)
	go handleInput(&camera, inputChannel)

	// spinningLine(screen)
	spinningCube(screen, &camera)
}

func spinningCube(screen Screen, camera *Camera) {
	camera.move(0.0, 0.0, -5.5)

	a := Point{1.0, 1.0, 1.0}
	b := Point{1.0, 1.0, -1.0}
	c := Point{1.0, -1.0, 1.0}
	d := Point{1.0, -1.0, -1.0}
	e := Point{-1.0, 1.0, 1.0}
	f := Point{-1.0, 1.0, -1.0}
	g := Point{-1.0, -1.0, 1.0}
	h := Point{-1.0, -1.0, -1.0}

	for theta := 0.0; theta < 1000; theta += 0.1 {
		screen.clear(dot)

		screen.drawLine(plot(e, camera), plot(a, camera), star)
		screen.drawLine(plot(a, camera), plot(b, camera), star)
		screen.drawLine(plot(b, camera), plot(f, camera), star)
		screen.drawLine(plot(f, camera), plot(e, camera), star)
		screen.drawLine(plot(g, camera), plot(c, camera), star)
		screen.drawLine(plot(c, camera), plot(d, camera), star)
		screen.drawLine(plot(d, camera), plot(h, camera), star)
		screen.drawLine(plot(h, camera), plot(g, camera), star)
		screen.drawLine(plot(a, camera), plot(c, camera), star)
		screen.drawLine(plot(b, camera), plot(d, camera), star)
		screen.drawLine(plot(f, camera), plot(h, camera), star)
		screen.drawLine(plot(e, camera), plot(g, camera), star)

		render(screen, green)

		time.Sleep(time.Millisecond * 100)
	}
}
