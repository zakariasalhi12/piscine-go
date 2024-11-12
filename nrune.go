package piscine

func NRune(s string, n int) rune {
	for lenght, v := range s {
		if lenght == n-1 {
			return v
		}
	}
	return 0
}
