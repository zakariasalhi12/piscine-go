package main

import "os"

func main() {
	if len(os.Args) == 1 {
		return
	}
	args := os.Args[1:]

	for _, elements := range args {
		if elements == "01" || elements == "galaxy 01" || elements == "galaxy" {
			os.Stdout.WriteString("Alert!!!")
		}
	}
}
