package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var t []bool
	x := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			x = false
			t = append(t, x)
		} else if int(f(a[i], a[i+1])) < 0 {
			x = true
			t = append(t, x)
		}
	}
	for j := 0; j < len(t)-1; j++ {
		if t[j] != t[j+1] {
			return false
		}
	}
	return true
}
