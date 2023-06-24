package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	i := 0
	for current != nil && i < pos {
		current = current.Next
		i++
	}
	if current == nil || i > pos {
		return nil
	}
	return current
}
