package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1
	for i := nb; i > 0; i-- {
		if result*i < result {
			return 0
		}
		result = result * i
	}

	return result
}
