package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

/*
Returns 0 if the strings are equal (s1==s2)
Returns 1 if string 1 is greater than string 2 (s1 > s2)
Returns -1 is string 1 is lesser than string 2 (s1 < s2)
*/
