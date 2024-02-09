package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		printhelp()
		return
	}
	args := os.Args[1:]

	order := false
	insert := ""
	res := ""

	if args[0] == "-h" || args[0] == "--help" {
		printhelp()
	}

	for _, element := range args {
		if element == "--order" || element == "-o" {
			order = true
		} else if contain(element, "--insert=") || contain(element, "-i=") {
			insert = element
		} else {
			res += element
		}
	}

	if insert != "" {
		if contain(insert, "insert=") {
			res += insert[9:]
		} else {
			res += insert[3:]
		}
	}

	if order {
		res = sortstring(res)
	}

	fmt.Print(res)

}

func sortstring(s string) string {
	slicerune := []rune(s)

	for i := 0; i < len(slicerune); i++ {
		for j := 0; j < len(slicerune); j++ {
			if slicerune[i] < slicerune[j] {
				slicerune[i], slicerune[j] = slicerune[j], slicerune[i]
			}
		}
	}
	return string(slicerune)
}

func contain(s, world string) bool {
	for i := 0; i <= len(s)-len(world); i++ {
		if s[i:i+len(world)] == world {
			return true
		}
	}
	return false
}

func printhelp() {
	fmt.Print("--insert\n  -i\n         This flag inserts the string into the string passed as argument.\n--order\n  -o\n         This flag will behave like a boolean, if it is called it will order the argument.")
}
