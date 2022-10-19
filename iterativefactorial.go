package piscine

func IterativeFactorial(nb int) int {
	result := 0
	if nb < 0 {
		result = 0
	} else if nb == 0 {
		result = 1
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
	
