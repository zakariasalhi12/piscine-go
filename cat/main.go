package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		copyContent(os.Stdin, os.Stdout)
	} else {
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				printError("ERROR: " + err.Error())
				os.Exit(1)
			}
			defer file.Close()

			copyContent(file, os.Stdout)
		}
	}
}

func copyContent(src io.Reader, dest io.Writer) {
	_, err := io.Copy(dest, src)
	if err != nil {
		printError("ERROR: " + err.Error())
		os.Exit(1)
	}
}

func printError(errMsg string) {
	for _, char := range errMsg {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
