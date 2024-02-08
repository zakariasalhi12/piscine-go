package piscine

func NRune(s string, n int) rune {
	for index, char := range s {
		if index+1 == n {
			return char
		}
	}
	return 0
}
