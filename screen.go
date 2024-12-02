package main

import (
	"fmt"
	"math"
)

type Screen [width][height]byte

func newScreen(char byte) Screen {
	var screen Screen
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			screen[x][y] = char
		}
	}
	return screen
}

func (s *Screen) clear(color byte) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			s[x][y] = color
		}
	}
}

func (s *Screen) drawDot(x, y int, color byte) {
	if x >= width {
		fmt.Println("Draw dot, x too large")
		fmt.Println(x)
		return
	}
	if x < 0 {
		fmt.Println("Draw dot, x below 0")
		fmt.Println(x)
		return
	}
	if y >= height {
		fmt.Println("Draw dot, y too large")
		fmt.Println(y)
		return
	}
	if y < 0 {
		fmt.Println("Draw dot, y below 0")
		fmt.Println(y)
		return
	}
	s[x][y] = color
}

func (s *Screen) drawLine(ax, ay, bx, by int, color byte) {
	dx := bx - ax
	if dx == 0 {
		for i := min(ay, by); i <= max(ay, by); i++ {
			s.drawDot(ax, i, color)
		}
		return
	}

	dy := by - ay
	if dy == 0 {
		for i := min(ax, bx); i <= max(ax, bx); i++ {
			s.drawDot(i, ay, color)
		}
		return
	}

	goingRight := dx > 0
	goingDown := dy > 0
	moreHorizontal := abs(dx) > abs(dy)

	var slope float64
	if moreHorizontal {
		slope = math.Abs(float64(dy)) / math.Abs(float64(dx))
	} else {
		slope = math.Abs(float64(dx)) / math.Abs(float64(dy))
	}

	vx := float64(ax)
	vy := float64(ay)
	s.drawDot(int(math.Round(vx)), int(math.Round(vy)), color)

	var increaseX float64
	var increaseY float64

	if moreHorizontal {
		if goingRight {
			increaseX = 1.0
		} else {
			increaseX = -1.0
		}
		if goingDown {
			increaseY = slope
		} else {
			increaseY = -slope
		}
	} else {
		if goingDown {
			increaseY = 1.0
		} else {
			increaseY = -1.0
		}
		if goingRight {
			increaseX = slope
		} else {
			increaseX = -slope
		}
	}

	for i := 0; i <= max(abs(dx), abs(dy)); i++ {
		vx += increaseX
		vy += increaseY
		s.drawDot(int(math.Round(vx)), int(math.Round(vy)), color)
	}

}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
