package piscine

func MakeRange(min, max int) []int {
	if max < min {
		return []int(nil)
	}
	arr := make([]int, max-min)
	for i := max - min; i < max; i++ {
		arr[i-min] = i
	}
	return arr
}
