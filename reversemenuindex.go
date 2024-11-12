package piscine

func ReverseMenuIndex(menu []string) []string {
	b := []string{}
	for i := len(menu) - 1; i >= 0; i-- {
		b = append(b, menu[i])
	}
	return b
}
