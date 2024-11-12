package piscine

func isPrime1(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for !isPrime1(nb) {
		nb++
	}
	return nb
}
