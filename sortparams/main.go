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

	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for _, elemnt := range args {
		for _, char := range elemnt {
			z01.PrintRune(char)
		}
		z01.PrintRune(10)
	}

}
