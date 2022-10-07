package piscine

import "github.com/01-edu/z01"

func FirstRune() {
	z01.PrintRune(Firstrune("Hello!"))
	z01.PrintRune(Firstrune("Salut!"))
	z01.PrintRune(Firstrune("Ola!"))
	z01.PrintRune('\n')
}

func Firstrune(s string) rune {
	r := rune(s[0])
	return r

}
