package piscine

func ToLower(s string) string {
	b := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			b += string(v + 32)
		} else {
			b += string(v)
		}
	}
	return b
}
