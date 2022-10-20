package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 25 {
		nb *= RecursiveFactorial(nb - 1)
	}
	return nb
}
