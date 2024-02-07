package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for m := j + 1; m <= '9'; m++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(m)
				if i != '7' || j != '8' || m != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
