package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string{}
	res := ""
	for _, char := range s {
		if char != ' ' {
			res += string(char)
		} else {
			if res != "" {
				arr = append(arr, res)
				res = ""
			}
		}
	}
	if res != "" {
		arr = append(arr, res)
	}
	return arr
}
