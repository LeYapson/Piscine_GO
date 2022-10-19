package piscine

func IterativeFactorial(nb int) int {
	a := nb
	var ap *int
	ap = &a
	if a < 0 {
		*ap = 0
	} else if a == 0 {
		*ap = 1
	} else {
		b := a - 1
		for i := b; i >= 1; i-- {
			*ap *= i
		}
	}
	return a
}
