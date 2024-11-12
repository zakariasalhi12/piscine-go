package piscine

func StrRev(s string) string {
	runes := string(s[len(s)-1])
	for i := len(s) - 2; i >= 0; i-- {
		runes += string(s[i])
	}
	return string(runes)
}
