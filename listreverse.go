package piscine

func ListReverse(l *List) {
	prev := nil
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
	l.Tail = curr
}
