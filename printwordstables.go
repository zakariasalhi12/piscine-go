package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	resault := ""
	for _, v := range a {
		resault += v + string('\n')
	}

	for _, r := range resault {
		z01.PrintRune(r)
	}
}
