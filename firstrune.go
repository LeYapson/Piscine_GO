package piscine


func FirstRune() {
	s := "trois"
	result := []rune(s) //convertis les string en runes
	FirstRune := string(result[0:1]) //convertis les runes de nouveau en string
println(FirstRune)
}