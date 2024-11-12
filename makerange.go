package piscine

func MakeRange(min, max int) []int {
	if min > max {
		return []int(nil)
	}

	if min == 0 && max == 0 {
		return []int(nil)
	}

	arr := make([]int, max-min)
	for i := 0; i < len2(arr); i++ {
		arr[i] = min + i
	}

	return arr
}

func len2(arr []int) int {
	counter := 0
	for range arr {
		counter++
	}
	return counter
}
