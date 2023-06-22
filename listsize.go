package piscine

func ListSize(l *List) int {
	count := 0
	temp := l.Head
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}
