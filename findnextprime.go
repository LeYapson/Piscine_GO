package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else if IsPrime(nb+1) {
		return nb + 1
	}
	return nb
}