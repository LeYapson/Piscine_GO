package piscine

func IterativePower(nb int, power int) int {
	if nb == 0 || power < 0 {
		return 0
	}
	a := 1
	for i := 0; i < power; i++ {
		a *= nb
	}
	return a
}