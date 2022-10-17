package piscine

func Capitalize(s string) string {
	my_array := []rune(s)
	for i, char := range my_array {
		if letterornumber(char) {
			if i == 0 || letterornumber(my_array[i-1]) == false {
				if my_array[i] >= 97 && my_array[i] <= 122 {
					my_array[i] = char - 32
				}
			} else {
				if my_array[i] >= 65 && my_array[i] <= 90 {
					my_array[i] = char + 32
				}
			}
		}
	}
	return string(my_array)
}

func letterornumber(r rune) bool {
	if r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122 {
		return true
	} else {
		return false
	}
}
