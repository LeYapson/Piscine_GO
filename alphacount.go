package piscine

func AlphaCount(s string) int {
	my_array := []rune(s)
	k := 0
	for _, r := range my_array { // for r := 0; r <= len(my_array-1); r++
		if r >= 65 && r <= 90 || r >= 97 && r <= 122 {
			k++
		}
	}
	return k
}

/*
- résulat : le nombre de caractères de l'alphabet qu'il y a au total
- 1 condition : ne doit pas compter les nombres
- 2 condition : ne dois pas compter les caracteres spéciaux
- 1: crée un tableau de rune
2 crée une variable int que totalise le nombre de caractère aqu'il y a



if r >= 65 && r <= 90(
	k++
) else if r >= 97 && r <= 122 (
	k++
)
)
return k

*/