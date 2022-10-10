package piscine

func Compare(a, b string) int {



if len(a) == len(b) {
	return 0
} else if len(a) > len(b){
	return 1
} else if len(a) < len(b) {
	return -1
}
return ('\n')
}


/*
Returns 0 if the strings are equal (s1==s2)
Returns 1 if string 1 is greater than string 2 (s1 > s2)
Returns -1 is string 1 is lesser than string 2 (s1 < s2)
*/