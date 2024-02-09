package piscine

func AppendRange(min, max int) []int {
	arr := []int(nil)
	if max < min {
		return arr
	}

	for i := max - min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
