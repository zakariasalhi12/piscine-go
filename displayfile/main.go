package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	var filename string

	if len(args) == 1 {
		filename = args[0]
	} else {
		fmt.Println("File name missing")
		return
	}

	content, _ := ioutil.ReadFile(filename)

	fmt.Print(string(content))
}
