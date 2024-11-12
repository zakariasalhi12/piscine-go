package piscine

func DescendAppendRange(max, min int) []int {
	arr := []int{}
	if max < min {
		return arr
	} else {
		for i := max; i > min; i-- {
			arr = append(arr, i)
		}
	}
	return arr
}
