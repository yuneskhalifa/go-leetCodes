package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	words := []string{}
	for _, char := range str {
		if char == ' ' {
			words = append(words, "")
		} else {
			words[len(words)-1] += string(char)
		}
	}
	for _, word := range words {
		summary[word]++
	}
	return summary
}
