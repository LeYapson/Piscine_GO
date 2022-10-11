package piscine

func IsNumeric(s string) bool {
	t := []rune(s)
	a := false
	for _, r := range t {
		if r >= 48 && r <= 57 {
			a = true
		} else {
			a = false
		}
	}
	return a
}
	