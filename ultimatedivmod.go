package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	rest := *a % *b
	*a = div
	*b = rest
}
