package piscine

func FindNextPrime(nb int) int {
	for i := nb+1; i > nb; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
