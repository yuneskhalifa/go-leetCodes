package piscine

func Index(s, toFind string) int {
	for i := 0; i < len(s); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}

	return -1
}
