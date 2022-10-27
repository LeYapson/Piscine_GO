package piscine

func Swap(a *int, b *int) {
	temp := 0
	temp = *a
	*a = *b
	*b = temp
}
