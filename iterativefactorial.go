package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	if nb < 0 {
		result = 0
	} else {
		for i := 1; i <= nb; i++ {
			result *= (i)
		}
	}
	return result
}
