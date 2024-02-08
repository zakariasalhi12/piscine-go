package piscine

import "github.com/01-edu/z01"

func isvalidebase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}

	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	res := ""
	negative := false

	if !isvalidebase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == 0 {
		z01.PrintRune('0')
		return
	}

	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbr(9)
		PrintNbr(223372036854775808)
		return
	}

	if nbr < 0 {
		negative = true
		nbr = -nbr
	}

	for nbr != 0 {
		res = string(base[nbr%len(base)]) + res
		nbr /= len(base)
	}

	if negative {
		res = "-" + res
	}

	for _, char := range res {
		z01.PrintRune(char)
	}
}
