package piscine

func Compact(ptr *[]string) int {
	pointer := []string(nil)
	for _, str := range *ptr {
		if str != "" {
			pointer = append(pointer, str)
		}
	}
	*ptr = pointer
	return len(pointer)
}
