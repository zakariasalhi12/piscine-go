package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1

	for _, char := range s {
		if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		} else if char == '+' && res == 0 {
			sign = 1

		} else if char == '-' && res == 0 {
			sign = -1
		}
	}

	return res * sign
}
