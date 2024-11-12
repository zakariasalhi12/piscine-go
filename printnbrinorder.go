package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	b := itoa(n)
	runes := []rune(b)

	for i := len(runes) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if runes[j] > runes[j+1] {
				temp := runes[j]
				runes[j] = runes[j+1]
				runes[j+1] = temp
			}
		}
	}

	showValue(runes)
}

func showValue(s []rune) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	for n != 0 {
		digit := n % 10
		result = string(rune(digit+'0')) + result
		n /= 10
	}

	return result
}
