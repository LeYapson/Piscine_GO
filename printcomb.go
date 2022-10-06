package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	
	for i := 48; i <= 57; i++ {
		a := rune(i)
		
		for j := 48; j <= 57; j++ {
			b := rune(j)
			
			for k := 48; k <= 57; k++ {
				c := rune(k)
				
				if i < j && j < k  {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
				}
			}	
		} 	

	} 						
}
