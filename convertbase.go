package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	res := 0
	index := 0
	if !isvalidebase(baseFrom) || !isvalidebase(baseTo) {
		return ""
	}

	for i := range nbr {
		for j := range baseFrom {
			if nbr[i] == baseFrom[j] {
				index = j
				break
			}
		}
		res = res*len(baseFrom) + index
	}

	reseul := ""
	for res != 0 {
		digit := res % len(baseTo)
		reseul = string(baseTo[digit]) + reseul
		res /= len(baseTo)
	}

	return reseul
}
