package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	inter := l.Tail.Data
	return inter
}
