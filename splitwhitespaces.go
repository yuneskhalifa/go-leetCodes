package piscine

func SplitWhiteSpaces(s string) []string {
	words := []string{}
	for _, rune := range s {
		if rune != ' ' && rune != '\t' && rune != '\n' {
			start := 0
			for _, rune := range s[start:] {
				if rune == ' ' || rune == '\t' || rune == '\n' {
					break
				}
				start++
			}
			words = append(words, s[start:start+1])
		}
	}
	return words
}
