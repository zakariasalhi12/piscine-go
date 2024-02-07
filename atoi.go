package piscine

func Atoi(s string) int {
	res := 0
	sign := 1
	for index, char := range s {
		if index == 0 && char == '-' {
			sign = -1
		} else if index == 0 && char == '+' {
			sign = 1
		} else if char >= '0' && char <= '9' {
			res = res*10 + int(char-48)
		} else {
			return 0
		}
	}

	return res * sign

}
