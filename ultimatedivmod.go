package piscine

func UltimateDivMod(a *int, b *int) {
	var div = *a / *b
	var rest = *a % *b
	*a = div
	*b = rest
}