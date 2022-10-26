package main

import (
	"piscine"

	"github.com/01-edu/z01"
)



type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	eachValuex := 0
	eachValuey := 0
	var arrayx []rune
	var arrayy []rune
	points := &point{}

	setPoint(points)

	for points.x != 0 {
		eachValuex = points.x % 10
		points.x /=10
		arrayx = append(arrayx, rune(eachValuex))
	}
	for points.y != 0 {
		eachValuey = points.y % 10
		points.y /=10
		arrayy = append(arrayy, rune(eachValuey))
	}
	text1 := "x = "
	text2 := ", y = "
	piscine.PrintStr(text1)
	for i:= len(arrayx) -1; i >= 0; i-- {
		z01.PrintRune(arrayx[i] + 48)
	}
	piscine.PrintStr(text2)
	for i:= len(arrayy) -1; i >= 0; i-- {
		z01.PrintRune(arrayy[i] + 48)
	}	
}
