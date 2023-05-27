package piscine

func SplitWhiteSpaces(s string) []string {
	words := []string{}
	for _, rune := range s {
		if rune == ' ' || rune == '\t' || rune == '\n' {
			continue
		}
		words = append(words, string(rune))
	}
	return words
}
