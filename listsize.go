package piscine

func ListSize(l *List) int {
	current := l.Head
	counter := 0

	for current != nil {
		current = current.Next
		counter++
	}

	return counter
}
