package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb >= 21 {
		return 0 // Overflow check for values greater than or equal to 21
	}

	result := nb * RecursiveFactorial(nb-1)

	return result
}
