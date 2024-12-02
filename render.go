package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func render(screen Screen) {
	buffer := ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			buffer += string(screen[x][y]) + " "
		}
		buffer += "\n"
	}
	clearScreen()
	fmt.Print(buffer)
}

func renderWithBorders(screen Screen) {
	buffer := ""
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
