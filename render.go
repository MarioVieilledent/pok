package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func render(screen Screen, color string) {
	buffer := color
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			buffer += string(screen[x][y]) + " "
		}
		buffer += "\n"
	}
	clearScreen()
	fmt.Print(buffer)
}

func renderWithBorders(screen Screen, color string) {
	buffer := color
	for x := 0; x < width+1; x++ {
		buffer += "--"
	}
	buffer += "\n"
	for y := 0; y < height; y++ {
		buffer += "|"
		for x := 0; x < width; x++ {
			buffer += string(screen[x][y]) + " "
		}
		buffer += "|\n"
	}
	for x := 0; x < width+1; x++ {
		buffer += "--"
	}
	clearScreen()
	fmt.Print(buffer)
}

func setTitle(title string) {
	fmt.Print("\x1b]0;" + title + "\x07")
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
