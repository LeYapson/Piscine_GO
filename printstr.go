package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	r := []rune(s)
	l := len(r)
	for i := 0; i < (l); i++ {
		z01.PrintRune(r[i])
	}
}
