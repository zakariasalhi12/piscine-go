package piscine

func IterativePower(nb int, power int) int {
	res := nb
	for i := 1; i < power; i++ {
		res *= nb
	}
	return res
}
