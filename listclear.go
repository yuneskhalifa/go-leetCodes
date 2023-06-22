package piscine

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		next := current.Next
		current = nil
		current = next
	}
}
