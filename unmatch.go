package piscine

func Unmatch(arr []int) int {
	for i, ch := range arr {
		count := 1
		for j, ch2 := range arr {
			if ch == ch2 && i != j {
				count++
			}
		}
		if count%2 == 1 {
			return ch
		}
	}
	return -1
}
