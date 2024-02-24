package piscine

func ActiveBits(n int) int {
	counter := 0
	for n != 0 {
		if n%2 == 1 {
			counter++
		}
		n = n / 2
	}
	return counter
}
