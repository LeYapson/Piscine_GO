package main

import "github.com/01-edu/z01"

type point struct {
	x 	int
	y 	int
	txt string
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
