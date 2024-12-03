package main

import (
	"log"

	"github.com/mattn/go-tty"
)

func listenInput(inputChannel chan string) {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		s, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		inputChannel <- string(s)
	}
}

func handleInput(camera *Camera, inputChannel chan string) {
	for key := range inputChannel {
		if key == "z" {
			camera.position[2] += speed
			camera.move(0, 0, speed)
		}
		if key == "q" {
			camera.position[0] -= speed
		}
		if key == "s" {
			camera.position[2] -= speed
		}
		if key == "d" {
			camera.position[0] += speed
		}
	}
}
