package piscine

func AlphaCount(s string) int {
	count := 0
	for _, s := range str {
		if isAlpha(s) {
			count++
		}
	}
	return count
}
