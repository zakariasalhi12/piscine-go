package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	args := os.Args[1:]
	upper := false

	for index, element := range args {
		elem := piscine.Atoi(element)
		if index == 0 && element == "--upper" {
			upper = true
		} else {
			if elem >= 1 && elem <= 26 {
				if upper {
					z01.PrintRune(rune(elem + 96 - 32))
				} else {
					z01.PrintRune(rune(elem + 96))
				}
			} else {
				z01.PrintRune(' ')
			}
		}

	}

}
