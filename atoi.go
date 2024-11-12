package piscine

func Atoi(s string) int {
	var result, sign int
	sign = 1

	for length, char := range s {
		switch {
		case char == '+':
			if length == 0 {
				sign = 1
			} else {
				return 0
			}
		case char == '-':
			if length == 0 {
				sign = -1
			} else {
				return 0
			}
		case char >= '0' && char <= '9':
			result = result*10 + int(char-'0')
		default:
			return 0
		}
	}

	return result * sign
}
