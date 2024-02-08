package piscine

func ToUpper(s string) string {
	res := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			res += string(char - 32)
		} else {
			res += string(char)
		}
	}
	return res
}
