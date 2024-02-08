package piscine

func RecursivePower(nb int, power int) int {
	if power == 1 {
		return nb
	}
	return nb * RecursivePower(nb, power-1)
}
