package piscine

func IterativeFactorial(nb int) int {
	result := 0
	if nb == 0 {
		return result
	} else {
		if nb >= 1 {
			result = 1
		for i := 1; i <= nb; i++ {
			result *= i
		}
		}
		
	}
	return result
}
