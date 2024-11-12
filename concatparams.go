package piscine

func ConcatParams(args []string) string {
	resault := ""
	for i, v := range args {
		if i != len(args)-1 {
			resault += v + "\n"
		} else {
			resault += v
		}
	}
	return resault
}
