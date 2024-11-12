package piscine

func BasicJoin(elems []string) string {
	b := ""

	for i := 0; i < len(elems); i++ {
		b += string(elems[i])
	}

	return string(b)
}
