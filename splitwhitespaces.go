package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	word := ""

	for _, char := range s {
		if char == ' ' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
