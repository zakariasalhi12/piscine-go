package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	negative := false

	res := ""
	if n == 0 {
		z01.PrintRune('0')
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
		z01.PrintRune('-')
	}

	for _, char := range sorted(res) {
		z01.PrintRune(char)
	}

}

func sorted(s string) string {
	res := []rune(s)
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] > res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return string(res)
}
