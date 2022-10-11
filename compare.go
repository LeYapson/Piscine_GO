package piscine

func Compare(a, b string) int {
	s1 := [] rune(a)
	s2 := [] rune(b)
	for i := 0; i < len(s1); i++ {
		if len(s1) == len(s2) {
			return 0
		} else if s1[0] != s2[0] {
			return -1
		} else if s1[len(s1)-1] != s2[len(s2)-1] {
			return 1
		}
	}
	return ('\n')
}


/*
Returns 0 if the strings are equal (s1==s2)
Returns 1 if string 1 is greater than string 2 (s1 > s2)
Returns -1 is string 1 is lesser than string 2 (s1 < s2)
*/
