package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	sor1 := true
	sor2 := true

	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) < 0 {
			sor2 = false
		}
	}

	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			sor1 = false
		}
	}

	return sor1 || sor2
}
