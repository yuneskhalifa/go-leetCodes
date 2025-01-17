package piscine

func IterativeFactorial(nb int) int {
	factorial := 1

	if nb >= 1 && nb <= 25 {
		for i := nb; i >= 1; i-- {
			factorial *= i
		}
		return factorial
	}
	if nb >= 26 || nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return 0
}
