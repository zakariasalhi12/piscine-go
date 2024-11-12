package piscine

func StringToIntSlice(str string) []int {
	resault := []int(nil)
	for _, char := range str {
		resault = append(resault, int(char))
	}
	return resault
}
