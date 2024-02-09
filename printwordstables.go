package piscine

func PrintWordsTables(a []string) {
	for index, elemnt := range a {
		if index != len(a)-1 {
			PrintStr(elemnt + "\n")
		} else {
			PrintStr(elemnt)
		}
	}
}
