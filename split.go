package piscine

func Split(s, sep string) []string {
	var result []string

	seplen := len(sep)
	start := 0

	for i := 0; i < len(s)-seplen; i++ {
		if s[i:i+seplen] == sep {
			result = append(result, s[start:i])
			start = i + seplen
		}
	}

	result = append(result, s[start:])
	return result
}
