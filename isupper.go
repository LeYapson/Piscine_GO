package piscine

func IsUpper(s string) bool {
	t := []rune(s)
	a := true
	for _, r := range t {
		if !(r >= 65 && r <= 90) {
			a = false
		}
	}
	return a
}
