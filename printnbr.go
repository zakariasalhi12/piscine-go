package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	res := ""
	negative := false

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n == -9223372036854775808 {
		z01.PrintRune('-')
		PrintNbr(9)
		PrintNbr(223372036854775808)
		return
	}

	if n < 0 {
		negative = true
		n = -n
	}

	for n != 0 {
		res = string(rune(n%10+'0')) + res
		n /= 10
	}

	if negative {
		res = "-" + res
	}

	for _, char := range res {
		z01.PrintRune(char)
	}
}
