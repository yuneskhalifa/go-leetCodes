package piscine

func Sqrt(nb int) int {
	z := 0
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			z = i
		}
	}
	return z
}
