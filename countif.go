package piscine

func CountIf(f func(string) bool, arr []string) int {
	resault := 0
	for _, res := range arr {
		if f(res) {
			resault += 1
		}
	}
	return resault
}
