package main

import (
	"fmt"
	"os"
)

func main() {
	b := os.Args[1:]
	for _, v := range b {
		if v == "01" || v == "galaxy" || v == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
