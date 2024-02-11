package main

import "fmt"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 41
	ptr.y = 20
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
