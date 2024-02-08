package piscine

func LastRune(s string) rune {
	result := []rune(s)
	return result[len(result)-1]
}
