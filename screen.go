package main

import (
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

func (s *Screen) clear(char byte) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			s[x][y] = char
		}
	}
}

func inBound(x, y int) bool {
	if x >= 0 && x < width && y >= 0 && y < height {
		return true
	}
	return false
}

func (s *Screen) drawPixel(pixel Pixel, char byte) {
	if inBound(pixel[0], pixel[1]) {
		s[pixel[0]][pixel[1]] = char
	}
}

func (s *Screen) drawLine(a, b Pixel, char byte) {
	dx := b[0] - a[0]
	if dx == 0 {
		for i := min(a[1], b[1]); i <= max(a[1], b[1]); i++ {
			s.drawPixel(Pixel{a[0], i}, char)
		}
		return
	}

	dy := b[1] - a[1]
	if dy == 0 {
		for i := min(a[0], b[0]); i <= max(a[0], b[0]); i++ {
			s.drawPixel(Pixel{i, a[1]}, char)
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

	vx := float64(a[0])
	vy := float64(a[1])
	s.drawPixel(Pixel{round(vx), round(vy)}, char)

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

	for i := 0; i < max(abs(dx), abs(dy)); i++ {
		vx += increaseX
		vy += increaseY
		s.drawPixel(Pixel{round(vx), round(vy)}, char)
	}
}

func round(v float64) int {
	return int(math.Round(v))
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
