package piscine

func ConcatParams(args []string) string {
	res := ""
	for index, element := range args {
		if index != len(args)-1 {
			res += element + "\n"
		} else {
			res += element
		}
	}
	return res
}
