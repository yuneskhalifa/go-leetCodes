package piscine

func Map(f func(int) bool, a []int) []bool {
	ad := []bool{}
	for _, elem := range a {
		ad = append(ad, f(elem))
	}
	return ad
}
