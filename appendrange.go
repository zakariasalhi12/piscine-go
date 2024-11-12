package piscine

func AppendRange(min, max int) []int {
	arr := []int(nil)
	if min > max {
		return arr
	} else {
		for i := min; i < max; i++ {
			arr = append(arr, i)
		}
	}
	return arr
}
