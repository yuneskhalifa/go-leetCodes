package piscine

func Map(f func(int) bool, a []int) []bool {
	ad := []bool{}
	for _, elem := range a {
		ad = append(a, f(elem))
	}
	return ad
}
