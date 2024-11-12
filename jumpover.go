package piscine

func JumpOver(str string) string {
	bara := ""
	for i := 2; i < len(str); i += 3 {
		bara += string(str[i])
	}
	bara += "\n"
	return bara
}
