package piscine

func StringToIntSlice(str string) []int {
	strin := []byte(str)
	res := []int(nil)
	for i := 0; i < len(strin); i++ {
		res = append(res, int(strin[i]))
	}
	return res
}
