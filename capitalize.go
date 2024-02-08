package piscine

func Capitalize(s string) string {
	result := ""
	firstChar := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if firstChar {
				if char >= 'a' && char <= 'z' {
					result += string(char - 32)
				} else {
					result += string(char)
				}
				firstChar = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char + 32)
				} else {
					result += string(char)
				}
			}
		} else {
			result += string(char)
			firstChar = true
		}
	}

	return result
}
