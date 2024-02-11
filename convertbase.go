package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	res := ""
	b := AtoiBase(nbr, baseFrom)

	if !isvalidebase(baseFrom) || !isvalidebase(baseTo) {
		return "NV"
	}
	res = PrintNbrBase2(b, baseTo)
	return res
}

func PrintNbrBase2(nbr int, base string) string {
	res := ""
	negative := false

	if nbr == 0 {
		return "0"
	}

	if nbr < 0 {
		negative = true
		nbr = -nbr
	}

	for nbr != 0 {
		res = string(base[nbr%len(base)]) + res
		nbr /= len(base)
	}

	if negative {
		res = "-" + res
	}

	return res
}
