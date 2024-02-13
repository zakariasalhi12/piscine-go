package piscine

func Rot14(s string) string {
	res := ""

	for _, char := range s {
		counter := char
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			for i := 0; i < 14; i++ {
				if counter == 'z' {
					counter = 'a'
				} else if counter == 'Z' {
					counter = 'A'
				} else {
					counter++
				}
			}
		}
		res = res + string(counter)
	}

	return res
}
