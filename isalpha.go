package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '4') {
			return false
		}
	}
	return true
}
