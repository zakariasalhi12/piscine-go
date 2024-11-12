package piscine

func Rot14(s string) string {
	result := ""
	for _, r := range s {
		counter := r

		if counter >= 'a' && counter <= 'z' {
			for i := 0; i < 14; i++ {
				counter++
				if counter > 'z' {
					counter = 'a'
				}
			}
			result += string(counter)
		} else if counter >= 'A' && counter <= 'Z' {
			for i := 0; i < 14; i++ {
				counter++
				if counter > 'Z' {
					counter = 'A'
				}
			}
			result += string(counter)
		} else {
			result += string(r)
		}
	}
	return result
}
