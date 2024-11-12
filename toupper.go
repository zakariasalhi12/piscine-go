package piscine

func ToUpper(s string) string {
	b := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			b += string(v - 32)
		} else {
			b += string(v)
		}
	}
	return b
}
