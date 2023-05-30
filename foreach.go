package piscine

func PrintNbr(b int) {
	Print(b)
}

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}
