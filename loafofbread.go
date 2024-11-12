package piscine

func LoafOfBread(str string) string {
	result := ""
	currentWord := ""

	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "Invalid Output\n"
	}

	for _, char := range str {
		if char != ' ' && len(currentWord) != 5 {
			currentWord += string(char)
		} else if len(currentWord) == 5 {
			if result != "" {
				result += " "
			}
			result += currentWord
			currentWord = ""
		}
	}
	if currentWord != "" {
		if result != "" {
			result += " "
		}
		result += currentWord
	}

	return result + "\n"
}
