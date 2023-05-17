package piscine

func UltimateDivMod(a *int, b *int) {
	x := *a / *b
	f := *a % *b
	*a = x
	*b = f
}
