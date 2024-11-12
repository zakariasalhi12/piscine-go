package piscine

func Join(strs []string, sep string) string {
	resault := ""
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			resault += strs[i] + sep
		} else {
			resault += strs[i]
		}
	}
	return resault
}
