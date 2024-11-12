package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			counter++
		}
	}
	return counter
}
