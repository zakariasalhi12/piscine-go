package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, compare func(a, b interface{}) bool) *interface{} {
	n := l.Head
	for n != nil {
		if compare(ref, n.Data) {
			return &n.Data
		}
		n = n.Next
	}
	return nil
}
