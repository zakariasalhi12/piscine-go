package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 || nb == 1 {
		return nb
	}

	for i := 2; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}

	return 0
}
