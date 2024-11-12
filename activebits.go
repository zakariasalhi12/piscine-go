package piscine

func ActiveBits(n int) int {
	counte := 0
	for n > 0 {
		if n%2 == 1 {
			counte++
		}

		n /= 2
	}

	return counte
}
