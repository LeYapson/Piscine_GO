package main

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
	txt := "x = %d, y = %d\n"
	points := &point{}

	setPoint(points)

	return txt
	return (points.x, points.y)
}