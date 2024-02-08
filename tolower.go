package piscine

func ToLower(s string) string {
	res := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			res += string(char + 32)
		} else {
			res += string(char)
		}
	}
	return res
}
