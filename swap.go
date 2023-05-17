package piscine

func Swap(a *int, b *int) {
	x := *a
	f := *b
	*a = f
	*b = x
}
