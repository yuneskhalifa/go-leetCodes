package piscine

func IterativeFactorial(nb int) int {
	factorial := 1

	for i := nb; i >= 1; i-- {
		factorial = factorial * i

	}
	return factorial

}
