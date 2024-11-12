package piscine

func BasicAtoi(s string) int {
	var result int

	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		}
	}

	return result
}
