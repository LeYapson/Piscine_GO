package piscine

func ToLower(s string) string {
	my_array := []rune(s)
	a := "\n"
	for _, i := range my_array {
		if (i >= 'A' && i <= 'Z') {
			my_array[i] -= 'A' + 'a'
		}
	}
	return a
}
