package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	var power float64 = 0.5
	var nbr float64 = float64(nb)
	nbr *= power
	if int(nbr)%2 != 0 {
		return 0
	} else if nb > 342939 {
		return 0
	}
	return int(nbr)
}