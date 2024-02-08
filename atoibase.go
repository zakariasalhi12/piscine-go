package piscine

func AtoiBase(s string, base string) int {
	res := 0
	index := 0
	if !isvalidebase(base) {
		return 0
	}

	for i := range s {
		for j := range base {
			if s[i] == base[j] {
				index = j
				break
			}
		}
		res = res*len(base) + index
	}
	return res
}
