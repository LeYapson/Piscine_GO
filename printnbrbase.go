package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	base_array := []rune(base)
for i :=0; i < len(base_array); i++ {
	if base_array < 2 {
		return "NV"
	} else if base_array[i] == base_array[i+1] {
		return "NV"
	} else if  base_array[i] == 43 && base_array[i] == 45 {
		return "NV"
	}
	if base_array >= 2 {
		if base_array[i] >= 48 && base_array[i] <= 57 {
			z01.PrintRune(nbr)
		} else if base_array[i] >= 48 && base_array[i] <= 49 {
			z01
		}
		
	}
}

}


/*
doit convertir le int dans la base donné par le str
si base = 0123456789 , doit ecrire 125
si base = 01 , doit ecrire 1111101
si base = 0123456789ABCDEF , doit écrire 7D
si base = choumi, doit écrire uoi
si base = aa , doit écrire NV