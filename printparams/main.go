package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	args := os.Args[1:]
	for _, elemnt := range args {
		for _, char := range elemnt {
			z01.PrintRune(char)
		}
		z01.PrintRune(10)
	}
}
