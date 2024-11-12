package piscine

func IsPrime(nb int) bool {
	if nb < 2 || nb <= 0 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
