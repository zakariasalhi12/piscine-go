package piscine

import (
	"github.com/01-edu/z01"
)

func prints(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func validbase(base string) bool {
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

func PrintNbrBase(n int, base string) {
	res := ""
	negative := false

	if !validbase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		negative = true
		n = n * -1
	}

	if n == -9223372036854775808 {
		prints("-")
		PrintNbrBase(9, base)
		PrintNbrBase(223372036854775808, base)
		return
	}

	for n != 0 {
		digits := n % len(base)

		res = string(base[digits]) + res

		n /= len(base)
	}

	if negative {
		res = "-" + res
		prints(res)
		return
	}

	prints(res)
}
