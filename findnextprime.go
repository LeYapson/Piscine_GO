package piscine

func FindNextPrime(nb int) int {
	for i := nb; i >= nb; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
