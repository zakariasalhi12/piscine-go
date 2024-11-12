package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	position := 0
	if len(os.Args) > 1 {
		if upper() {
			for i := 2; i < len(os.Args); i++ {
				position = atoiha(os.Args[i])
				if position >= 1 && position <= 26 {
					if upper() {
						z01.PrintRune(rune(position + 96 - 32))
					} else {
						z01.PrintRune(rune(position + 96))
					}
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		} else {
			for i := 1; i < len(os.Args); i++ {
				position = atoiha(os.Args[i])
				if position >= 1 && position <= 26 {
					if upper() {
						z01.PrintRune(rune(position + 96 - 32))
					} else {
						z01.PrintRune(rune(position + 96))
					}
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func upper() bool {
	if os.Args[1] == "--upper" {
		return true
	}
	return false
}

func atoiha(s string) int {
	var result int

	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 0
		}
	}

	return result
}
