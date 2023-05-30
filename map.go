package piscine

func Map(f func(int) bool, a []int) []bool {
	ad := []bool{}
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			ad = append(ad, true)
		} else {
			ad = append(ad, false)
		}
	}
	return ad
}
