package piscine

func Sqrt(nb int) int {
	z := 0
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			z = 1
		}
	}
	return z
}
