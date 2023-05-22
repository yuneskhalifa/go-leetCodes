package piscine

func AlphaCount(s string) int {
	count := 0
	for _ , a := range s {
		if isAlpha(a) {
			count++
		}
	}
	return count
}
