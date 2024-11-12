package main

import (
	"github.com/01-edu/z01"
)

const m = "x = 42, y = 21"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	for _, v := range m {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
