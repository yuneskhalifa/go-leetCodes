package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a); i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	for j := 0; j < len(a); j++ {
		if f(a[j], a[j+1]) < 0 {
			return true
		}
	}
	return false
}
